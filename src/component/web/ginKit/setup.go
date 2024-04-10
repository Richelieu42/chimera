package ginKit

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
	"github.com/sirupsen/logrus"
	"time"
)

// serviceInfo e.g."Agent-127.0.0.1:12345"
var serviceInfo = ""

func MustSetUp(config *Config, businessLogic func(engine *gin.Engine) error, options ...GinOption) {
	err := SetUp(config, businessLogic, options...)
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// SetUp
/*
PS: 正常执行的情况下，此方法会阻塞调用的协程.

@param config			不能为nil（否则将返回error）
@param businessLogic 	业务逻辑，可以在其中进行 路由绑定 等操作...（可以为nil，但不推荐这么干）
*/
func SetUp(config *Config, businessLogic func(engine *gin.Engine) error, options ...GinOption) error {
	if err := validateKit.Struct(config); err != nil {
		return err
	}

	opts := loadOptions(options...)
	serviceInfo = opts.ServiceInfo

	// Gin的模式，后续也可以在 businessLogic 里面调整
	if strKit.IsEmpty(config.Mode) {
		// 默认: debug模式
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(config.Mode)
	}

	/*
		gin框架中如何让日志文字带颜色输出？
			https://mp.weixin.qq.com/s/eHtIC5egDoqx4LdAvcE5Qw

		PS: 如果 Gin.ForceConsoleColor() 和 Gin.DisableConsoleColor() 都不调用，那么默认是在终端中输出日志是带颜色的，输出到其他地方是不带颜色的.
	*/
	if config.DisableColor {
		// 禁用日志带颜色输出
		gin.DisableConsoleColor()
	} else {
		// 强制日志带颜色输出（无论是在终端还是其他输出设备）
		gin.ForceConsoleColor()
	}

	// 通过logrus输出Gin的日志.
	// Richelieu：从目前表现来看，虽然gin和logrus都可以设置颜色，但在此处,只要gin允许了，logrus的logger是否允许就无效了
	gin.DefaultWriter = logrus.StandardLogger().Out

	engine := NewEngine()
	/*
		MaxMultipartMemory只是限制内存，不是针对文件上传文件大小，即使文件大小比这个大，也会写入临时文件。
		默认32MiB，并不涉及"限制上传文件的大小"，原因：上传的文件s按顺序存入内存中，累加大小不得超出 32Mb ，最后累加超出的文件就存入系统的临时文件中。非文件字段部分不计入累加。所以这种情况，文件上传是没有任何限制的。
		参考: https://studygolang.com/articles/22643
	*/
	engine.MaxMultipartMemory = 64 << 20

	// pprof
	if config.Pprof {
		pprof.Register(engine, pprof.DefaultPrefix) // 等价于 pprof.Register(engine)
	}

	// middleware
	if err := attachMiddlewares(engine, config.Middleware, opts); err != nil {
		return err
	}

	/* favicon.ico */
	if opts.DefaultFavicon {
		DefaultFavicon(engine)
	}

	/* 404 */
	if opts.DefaultNoRouteHtml {
		if err := DefaultNoRouteHtml(engine); err != nil {
			return err
		}
	}

	/* 405（不设置的话，就会走到404） */
	engine.HandleMethodNotAllowed = opts.DefaultNoMethod
	if engine.HandleMethodNotAllowed {
		DefaultNoMethod(engine)
	}

	/* 业务逻辑 */
	if businessLogic != nil {
		if err := businessLogic(engine); err != nil {
			return errorKit.Wrapf(err, "Fail to execute businessLogic().")
		}
	}

	ssl := config.SSL
	if ssl.Port != 0 {
		if config.Port == 0 {
			// (1) https port（本服务使用1个端口）
			if err := netKit.AssertValidPort(ssl.Port); err != nil {
				return errorKit.Wrapf(err, "Https port(%d) should be set to a valid value.", ssl.Port)
			}
			return engine.RunTLS(netKit.JoinHostnameAndPort(config.HostName, ssl.Port), ssl.CertFile, ssl.KeyFile)
		}
		// (2) https port + http port（本服务使用2个端口）
		if err := netKit.AssertValidPort(ssl.Port); err != nil {
			return errorKit.Wrapf(err, "Https port(%d) should be set to a valid value.", ssl.Port)
		}
		if err := netKit.AssertValidPort(config.Port); err != nil {
			return errorKit.Wrapf(err, "Http port(%d) should be set to a valid value.", config.Port)
		}
		if config.Port == ssl.Port {
			return errorKit.Newf("Http port and https port are same(%d).", config.Port)
		}

		go func() {
			if err := engine.Run(netKit.JoinHostnameAndPort(config.HostName, config.Port)); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{
					"port": config.Port,
				}).Fatalf("[%s, GIN] Fail to start http server.", consts.UpperProjectName)
			}
		}()
		go func() {
			time.Sleep(time.Millisecond * 20)

			if err := engine.RunTLS(netKit.JoinHostnameAndPort(config.HostName, ssl.Port), ssl.CertFile, ssl.KeyFile); err != nil {
				logrus.WithError(err).WithFields(logrus.Fields{
					"port":     ssl.Port,
					"certFile": ssl.CertFile,
					"keyFile":  ssl.KeyFile,
				}).Fatalf("[%s, GIN] Fail to start https server.", consts.UpperProjectName)
			}
		}()
		select {}
	}
	// (3) http port（本服务使用1个端口）
	if config.Port == 0 {
		return errorKit.Newf("At least one of http port and https port should be set to a valid value.")
	}
	if err := netKit.AssertValidPort(config.Port); err != nil {
		return errorKit.Wrapf(err, "Http port(%d) should be set to a valid value.", config.Port)
	}
	return engine.Run(netKit.JoinHostnameAndPort(config.HostName, config.Port))
}
