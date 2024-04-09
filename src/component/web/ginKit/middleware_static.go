package ginKit

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

// NewStaticMiddleware
/*
PS:
(1) 可以与已绑定的路由冲突且优先级高;
(2) 适用场景: 静态资源与已绑定路由冲突的情况（如果不冲突的话，建议使用 StaticDir）.

@param root		目录路径
@param indexes 	是否列出目录信息（安全考虑的话，建议传参false）
*/
func NewStaticMiddleware(urlPrefix string, dirPath string, indexes bool) (gin.HandlerFunc, error) {
	if err := fileKit.AssertExistAndIsDir(dirPath); err != nil {
		return nil, err
	}

	fs := static.LocalFile(dirPath, indexes)
	return static.Serve(urlPrefix, fs), nil
}

func NewStaticMiddlewareWithEmbedFolder(urlPrefix string, fsEmbed embed.FS, targetPath string) gin.HandlerFunc {
	fs := static.EmbedFolder(fsEmbed, targetPath)
	return static.Serve(urlPrefix, fs)
}
