package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"net/http"
)

// StaticFile 静态资源（单个文件）
/*
@param relativePath	路由
@param filePath 	相对路径（对于项目的根目录(working directory)，而非main()所在的目录（虽然他们常常是同一个）） || 绝对路径
*/
func StaticFile(group gin.IRoutes, relativePath, filePath string) error {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return err
	}

	group.StaticFile(relativePath, filePath)
	return nil
}

// StaticDir 静态资源（目录）
/*
PS: 适用场景: 静态资源与已绑定路由不冲突的情况（如果冲突的话，建议使用 NewStaticMiddleware）.

@param relativePath		路由
@param root				相对路径（对于项目的根目录(working directory)，而非main()所在的目录（虽然他们常常是同一个）） || 绝对路径
@param listDirectory 	是否列出目录下的文件，true: 当目录下不存 index.html 文件时，会列出该目录下的所有文件（正式环境不推荐，因为不安全）
*/
func StaticDir(group gin.IRoutes, relativePath, root string, listDirectory bool) error {
	if err := fileKit.AssertExistAndIsDir(root); err != nil {
		return err
	}

	fs := gin.Dir(root, listDirectory)
	return StaticFS(group, relativePath, fs)
}

func StaticFS(group gin.IRoutes, relativePath string, httpFs http.FileSystem) error {
	if err := interfaceKit.AssertNotNil(httpFs, "httpFs"); err != nil {
		return err
	}

	group.StaticFS(relativePath, httpFs)
	return nil
}
