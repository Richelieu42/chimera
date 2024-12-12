package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, "Hello world!")
	})

	if err := engine.RunTLS(":443", "ssl.crt", "ssl.key"); err != nil {
		console.Fatalf("Fail to run, error: %s", err)
	}
}
