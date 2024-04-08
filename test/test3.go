package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, strings.Repeat("c", 2000))
	})

	if err := engine.Run(":8888"); err != nil {
		logrus.Fatal(err)
	}
}
