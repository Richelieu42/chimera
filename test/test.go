package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/proxy/forwardKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"net/url"
)

func main() {
	port := 80

	u, err := url.Parse("http://127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	rp, err := forwardKit.NewSingleHostReverseProxy(u)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {

		ctx.String(200, fmt.Sprintf("This is [%d].", port))
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
