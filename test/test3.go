package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/proxyKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/wsKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

var target = "127.0.0.1:8000"

func main() {
	engine := gin.Default()

	ginKit.BindHandlersToRoute(engine, "/connection/websocket", []string{http.MethodPost, http.MethodGet}, func(ctx *gin.Context) {
		wsKit.PolyfillWebSocketRequest(ctx.Request)

		proxyCF(ctx)
	})
	ginKit.BindHandlersToRoutes(engine, []string{"/connection/sse", "/connection/http_stream", "/emulation"}, []string{http.MethodPost, http.MethodGet}, proxyCF)

	if err := engine.Run(":80"); err != nil {
		logrus.Fatal(err)
	}
}

func proxyCF(ctx *gin.Context) {
	if err := proxyKit.ProxyWithGin(ctx, target, proxyKit.WithErrorLogger(nil)); err != nil {
		logrus.WithError(err).WithField("route", httpKit.GetRoute(ctx.Request)).Error("Fail to proxy.")
	}
}
