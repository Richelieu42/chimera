package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, "8002")
	})

	if err := engine.Run(":8002"); err != nil {
		panic(err)
	}
}
