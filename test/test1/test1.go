package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		logrus.Infof("Host: [%s]", ctx.Request.Host)
		logrus.Infof("URL.Host: [%s]", ctx.Request.URL.Host)

		ctx.String(200, "8001")
	})

	if err := engine.Run(":8001"); err != nil {
		panic(err)
	}
}
