package ginKit

import "github.com/gin-gonic/gin"

// LoadHTMLFiles 加载（多个）html文件
/*
Deprecated: 直接调用 IEngine 的方法.
*/
func LoadHTMLFiles(engine *gin.Engine, filePaths ...string) {
	engine.LoadHTMLFiles(filePaths...)
}

// LoadHTMLGlob
/*
Deprecated: 直接调用 IEngine 的方法.
*/
func LoadHTMLGlob(engine *gin.Engine, pattern string) {
	engine.LoadHTMLGlob(pattern)
}
