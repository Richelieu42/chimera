package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"time"
)

func main() {
	port := 8001

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		zapKit.Info("get a request")

		time.Sleep(time.Second * 5)

		ctx.String(200, fmt.Sprintf("This is [%d].", port))
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
