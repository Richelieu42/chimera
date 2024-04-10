package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/internal/resources"
)

func DefaultFavicon(engine *gin.Engine) {
	path := "_resources/icon/favicon.ico"
	fs := resources.AssetFile()

	//httpFileSystem := http.FS(efs)

	engine.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.FileFromFS(path, fs)

		//SetResponseHeader(ctx, "Cache-Control", "public, max-age=31536000")
		//ctx.FileFromFS("_icon/favicon.ico", httpFileSystem)
	})
}
