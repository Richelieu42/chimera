package ginKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultFavicon(engine *gin.Engine) {
	httpFileSystem := http.FS(efs)

	engine.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.FileFromFS("_icon/favicon.ico", httpFileSystem)
	})
}
