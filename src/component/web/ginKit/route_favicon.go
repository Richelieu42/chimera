package ginKit

import (
	"github.com/gin-gonic/gin"
)

func DefaultFavicon(engine *gin.Engine) {
	//path := "_resources/icon/favicon.ico"
	//fs := resources.AssetFile()

	engine.GET("/favicon.ico", func(ctx *gin.Context) {

		efs.Open("_icon/favicon.ico")

		//ctx.FileFromFS(path, fs)
	})
}
