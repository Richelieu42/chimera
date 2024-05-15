package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/proxyKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/wsKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

var target = "127.0.0.1:8000"

func main() {
	engine := gin.Default()

	// gzip中间件
	engine.Use(ginKit.NewGzipMiddleware(1))
	//engine.Use(ginKit.NewGzipMiddleware(1, gzip.WithExcludedPaths([]string{"/connection/http_stream"})))
	//engine.Use(ginKit.NewGzipMiddleware2(1, 0))

	// cors中间件
	engine.Use(ginKit.NewCorsMiddleware(nil))

	// options中间件
	engine.Use(ginKit.NewOptionsMiddleware())

	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, "heLLo")
	})
	ginKit.BindHandlersToRoute(engine, "/connection/websocket", []string{http.MethodGet}, func(ctx *gin.Context) {
		wsKit.PolyfillWebSocketRequest(ctx.Request)

		proxyToCentrifugo(ctx)
	})
	ginKit.BindHandlersToRoute(engine, "/connection/sse", []string{http.MethodGet}, proxyToCentrifugo)
	ginKit.BindHandlersToRoutes(engine, []string{"/connection/http_stream", "/emulation"}, []string{http.MethodPost}, func(ctx *gin.Context) {
		proxyToCentrifugo(ctx)
	})

	if err := engine.Run(":80"); err != nil {
		logrus.Fatal(err)
	}
}

func proxyToCentrifugo(ctx *gin.Context) {
	// Richelieu: 删掉允许跨域头，以防双重允许跨域（centrifugo服务那边已经有允许跨域了）
	httpKit.DelHeader(ctx.Writer.Header(), httpKit.HeaderAccessControlAllowOrigin)

	if err := proxyKit.ProxyWithGin(ctx, target, proxyKit.WithErrorLogger(nil), proxyKit.WithModifyResponse(modifyResponse)); err != nil && !errors.Is(err, http.ErrAbortHandler) {
		logrus.WithError(err).WithField("route", httpKit.GetRoute(ctx.Request)).Error("Fail to proxy.")
		return
	}
}

func modifyResponse(response *http.Response) error {
	// 修改 Connection 头为 "keep-alive"
	response.Header.Set("Connection", "keep-alive")
	return nil
}
