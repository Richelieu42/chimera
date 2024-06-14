package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/proxy/forwardKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

/*
访问地址: http://127.0.0.1:8000/test
*/
func main() {
	/* 目标服务(target) */
	go func() {
		port := 8001

		engine := gin.Default()
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(555, fmt.Sprintf("This is [%d].", port))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	/* 代理服务 */
	errLog := logKit.NewStdoutLogger("")
	engine := gin.Default()
	var modifyResponse func(resp *http.Response) error
	modifyResponse = func(resp *http.Response) error {
		if resp.StatusCode != 200 {
			return errorKit.Simplef("invalid response with status(%s)", resp.Status)
		}
		return nil
	}
	//modifyResponse = nil
	engine.Any("/test", func(ctx *gin.Context) {
		err := forwardKit.ForwardToHostComplexly(ctx.Writer, ctx.Request, "127.0.0.1:8001", errLog, nil, modifyResponse)
		if err != nil {
			logrus.WithError(err).Info("Fail to forward.")
			ctx.String(http.StatusBadGateway, err.Error())
			return
		}
		logrus.Info("Manager to forward.")
		return
	})
	if err := engine.Run(":8000"); err != nil {
		panic(err)
	}
}
