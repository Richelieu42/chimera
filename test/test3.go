package main

import (
	"errors"
	"github.com/gin-contrib/gzip"
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/wsKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"

	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/proxyKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

var target = "127.0.0.1:8000"

func main() {
	engine := gin.Default()

	// gzip压缩
	engine.Use(ginKit.NewGzipMiddleware(1, gzip.WithExcludedPaths([]string{"/connection/http_stream"})))
	// 允许跨域
	engine.Use(ginKit.NewCorsMiddleware(nil))

	ginKit.BindHandlersToRoute(engine, "/connection/websocket", []string{http.MethodGet}, func(ctx *gin.Context) {
		wsKit.PolyfillWebSocketRequest(ctx.Request)

		proxyToCentrifugo(ctx)
	})
	ginKit.BindHandlersToRoutes(engine, []string{"/connection/sse"}, []string{http.MethodGet}, proxyToCentrifugo)
	ginKit.BindHandlersToRoutes(engine, []string{"/connection/http_stream", "/emulation"}, []string{http.MethodPost, http.MethodOptions}, proxyToCentrifugo)

	if err := engine.Run(":80"); err != nil {
		logrus.Fatal(err)
	}
}

func proxyToCentrifugo(ctx *gin.Context) {
	if err := proxyKit.ProxyWithGin(ctx, target, proxyKit.WithErrorLogger(nil)); err != nil && !errors.Is(err, http.ErrAbortHandler) {
		logrus.WithError(err).WithField("route", httpKit.GetRoute(ctx.Request)).Error("Fail to proxy.")
		return
	}
}
