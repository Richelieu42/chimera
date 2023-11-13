// Package ginKit
// 静态资源相关
package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

// LoadHTMLFiles 加载（多个）html文件
/*
Deprecated: 直接调用 IEngine 的方法.
*/
func LoadHTMLFiles(engine IEngine, filePaths ...string) {
	engine.LoadHTMLFiles(filePaths...)
}

// LoadHTMLGlob
/*
Deprecated: 直接调用 IEngine 的方法.
*/
func LoadHTMLGlob(engine IEngine, pattern string) {
	engine.LoadHTMLGlob(pattern)
}

// StaticFile 静态资源（单个文件）
/*
@param relativePath 路由
@param filePath 相对路径（对于项目的根目录(working directory)，而非main()所在的目录（虽然他们常常是同一个）） || 绝对路径
*/
func StaticFile(group IGroup, relativePath, filePath string) error {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return err
	}

	group.StaticFile(relativePath, filePath)
	return nil
}

// StaticDir 静态资源（目录）
/*
@param relativePath	路由
@param root 相对路径（对于项目的根目录(working directory)，而非main()所在的目录（虽然他们常常是同一个）） || 绝对路径
*/
func StaticDir(group IGroup, relativePath, root string, listDirectory bool) error {
	if err := fileKit.AssertExistAndIsDir(root); err != nil {
		return err
	}

	fs := gin.Dir(root, listDirectory)
	group.StaticFS(relativePath, fs)
	return nil
}
