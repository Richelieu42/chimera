package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	_ "github.com/richelieu-yang/chimera/v3/src/component/mq/pulsarKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	_ "github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	_ "github.com/richelieu-yang/chimera/v3/src/statKit"
	_ "go.uber.org/automaxprocs"
)

var port int

func init() {
	flag.IntVar(&port, "port", 80, "address")
}

func main() {
	flag.Parse()

	console.PrintBasicDetails()

	c := &ginKit.Config{
		Port:         port,
		DisableColor: false,
		Pprof:        true,
		SSL:          ginKit.SslConfig{},
		Middleware:   ginKit.MiddlewareConfig{},
	}

	ginKit.MustSetUp(c, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(200, "ok")
			return
		})

		return nil
	}, ginKit.WithServiceInfo("TEST"), ginKit.WithDefaultFavicon(true), ginKit.WithDefaultNoRouteHtml(true))
}
