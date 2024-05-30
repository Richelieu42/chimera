package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	engine := ginKit.NewEngine()

	engine.Any("/test", func(ctx *gin.Context) {
		select {
		case <-ctx.Done():
			logrus.Warn("ctx.Done()")
			return
		case <-time.After(time.Second * 10):
			ctx.String(200, "This is 12000.")
		}
	})

	if err := engine.Run(":12000"); err != nil {
		panic(err)
	}
}
