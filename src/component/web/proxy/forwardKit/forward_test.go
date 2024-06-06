package forwardKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/log/logKit"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"testing"
)

func TestForwardToUrl(t *testing.T) {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		errLogger := logKit.NewLogger(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
		err := ForwardToUrl(ctx.Writer, ctx.Request, errLogger, "http://localhost:20000/test")
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
