package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
)

func main() {
	type config struct {
		Gin *ginKit.Config `json:"gin" yaml:"gin"`
	}

	path := "_chimera-lib/config.yaml"
	c := &config{}
	//err := yamlKit.UnmarshalFromFile(path, c)
	_, err := viperKit.UnmarshalFromFile(path, nil, c)
	if err != nil {
		panic(err)
	}

	ginKit.MustSetUp(c.Gin, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(200, "ok")
			return
		})

		return nil
	}, ginKit.WithServiceInfo("TEST"), ginKit.WithDefaultFavicon(true), ginKit.WithDefaultNoRouteHtml(true))
}
