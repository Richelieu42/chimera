package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"

	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/netKit"

	"github.com/gin-gonic/gin"
)

func main() {
	port := 8002

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, fmt.Sprintf("This is [%d].", port))
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
