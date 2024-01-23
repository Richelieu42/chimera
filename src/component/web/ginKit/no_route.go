package ginKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/internal/resources"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"net/http"
	"sync"
)

var noRouteOnce sync.Once
var noRouteErr error

// NoRoute 404
func NoRoute(engine IEngine, handlers ...gin.HandlerFunc) {
	engine.NoRoute(handlers...)
}

// DefaultNoRouteHtml 使用自带的404页面.
func DefaultNoRouteHtml(engine IEngine) error {
	htmlPath := "_resources/html/404.min.html"

	noRouteOnce.Do(func() {
		tempDir, err := pathKit.GetExclusiveTempDir()
		if err != nil {
			noRouteErr = err
			return
		}
		err = resources.RestoreAsset(tempDir, htmlPath)
		filePath := pathKit.Join(tempDir, htmlPath)
		if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
			noRouteErr = err
			return
		}
		engine.LoadHTMLFiles(filePath)
	})
	if noRouteErr != nil {
		return noRouteErr
	}

	var prefix string
	if strKit.IsNotEmpty(serviceInfo) {
		prefix = fmt.Sprintf("[%s] ", serviceInfo)
	}
	name := fileKit.GetFileName(htmlPath)
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, name, gin.H{
			"prefix": prefix,
			"route":  ctx.Request.URL.Path,
		})

		//ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", noRouteData)

		//// 此处不使用 ctx.FileFromFS ，原因: 这样的话，响应状态码就会是200，改不了
		//fs := resources.AssetFile()
		//ctx.FileFromFS("_resources/html/404.min.html", fs)
	})
	return nil
}
