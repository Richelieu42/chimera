package ginKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultFavicon(engine *gin.Engine) {
	//path := "_resources/icon/favicon.ico"
	//fs := resources.AssetFile()

	engine.GET("/favicon.ico", func(ctx *gin.Context) {
		//ctx.FileFromFS(path, fs)

		ctx.FileFromFS("_icon/favicon.ico", http.FS(efs))
	})
}
