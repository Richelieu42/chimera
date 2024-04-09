package ginKit

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// NewStaticMiddleware
/*
@param root
@param indexes
*/
func NewStaticMiddleware(urlPrefix string, root string, indexes bool) gin.HandlerFunc {
	fs := static.LocalFile(root, indexes)
	return static.Serve(urlPrefix, fs)
}

func NewStaticMiddlewareWithEmbedFolder(urlPrefix string, fsEmbed embed.FS, targetPath string) gin.HandlerFunc {
	fs := static.EmbedFolder(fsEmbed, targetPath)
	return static.Serve(urlPrefix, fs)
}
