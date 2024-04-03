package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/internal/resources"
)

func DefaultFavicon(engine *gin.Engine) {
	path := "_resources/icon/favicon.ico"
	fs := resources.AssetFile()

	engine.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.FileFromFS(path, fs)
	})
}
