package proxyKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestProxy(t *testing.T) {
	engine := gin.Default()

	engine.Any("/proxy", func(ctx *gin.Context) {
		err := Proxy(ctx.Writer, ctx.Request, "127.0.0.1:8888")
		if err != nil {
			if IsProxyDialError(err) {
				ctx.String(520, "proxy dial error")
				return
			}

			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.String(http.StatusOK, "OK")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
