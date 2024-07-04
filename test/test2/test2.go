package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
)

func main() {
	port := 8001

	engine := gin.Default()
	engine.POST("/test", func(ctx *gin.Context) {
		ctx.String(200, "Hello world!")
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
