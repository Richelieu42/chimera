package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
)

func main() {
	middleware, err := ginKit.NewStaticMiddleware("/", "_chimera-lib", true)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	engine.Use(middleware)

	group := engine.Group("g")
	//group.Use(middleware)
	group.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
