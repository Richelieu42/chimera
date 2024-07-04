package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
)

func main() {
	port := 8001

	engine := gin.Default()
	engine.POST("/test", func(ctx *gin.Context) {
		name := ginKit.ObtainPostParam(ctx, "name")
		age := ginKit.ObtainPostParam(ctx, "age")
		ctx.String(200, fmt.Sprintf("Hello %s(%s)", name, age))
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
