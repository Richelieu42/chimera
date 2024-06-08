package forwardKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/log/logKit"
	"github.com/sirupsen/logrus"
	"testing"
)

/*
访问url: http://127.0.0.1/test
效果: 	将 http://127.0.0.1/test 转发给 http://127.0.0.1:8000/test
*/
func TestForwardToUrl(t *testing.T) {
	url := "http://127.0.0.1:8000"

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		errLog := logKit.NewStdoutLogger("")
		err := ForwardToUrl(ctx.Writer, ctx.Request, url, errLog)
		if err != nil {
			logrus.WithError(err).Info("Fail to forward.")
			ctx.String(500, err.Error())
			return
		}
		logrus.Info("Manager to forward.")
		return
	})
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
