package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"go.uber.org/zap"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		console.Panic("len(os.Args) invalid", zap.Int("length", len(os.Args)))
	}
	certFile := os.Args[1]
	keyFile := os.Args[2]

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, "Hello world!")
	})
	if err := engine.RunTLS(":666", certFile, keyFile); err != nil {
		console.Fatalf("Fail to run, error: %s", err)
	}
}
