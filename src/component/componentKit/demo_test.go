package componentKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/pathKit"
	"testing"
)

func Test(t *testing.T) {
	if err := pathKit.ReviseProjectDirWhenTesting(); err != nil {
		panic(err)
	}

	if err := InitializeEnvironment(); err != nil {
		errorKit.Panic("fail to initialize %s, error: %+v", "env", err)
	}

	// redis组件（可选）
	if err := InitializeRedisComponent(); err != nil {
		errorKit.Panic("fail to initialize %s, error: %+v", "redis", err)
	}

	if err := business(); err != nil {
		errorKit.PanicByError(err)
	}

	// json组件（可选）
	msgProcessor := func(code string, msg string, data interface{}) string {
		// TODO: 额外处理message
		return msg
	}
	if err := InitializeJsonComponent(msgProcessor, ""); err != nil {
		errorKit.Panic("fail to initialize %s, error: %+v", "json", err)
	}

	// gin组件（可选）
	recoveryMiddleware := gin.CustomRecovery(func(c *gin.Context, err any) {
		// TODO: gin处理请求时发生panic的情况，在此处进行相应的处理（比如响应json给前端）
	})
	if err := InitializeGinComponent(recoveryMiddleware, routeBusiness); err != nil {
		errorKit.Panic("fail to initialize %s, error: %+v", "gin", err)
	}
}

func business() error {
	// TODO: 业务逻辑（读取业务配置文件...）

	return nil
}

func routeBusiness(engine *gin.Engine) error {
	// TODO: 业务逻辑（绑定路由...）

	return nil
}
