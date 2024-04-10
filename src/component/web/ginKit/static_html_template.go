package ginKit

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

// LoadHTMLFiles 加载（多个）html文件.
/*
PS: 搭配 gin.Context.HTML() 使用.
*/
func LoadHTMLFiles(engine *gin.Engine, filePaths ...string) {
	engine.LoadHTMLFiles(filePaths...)
}

// LoadHTMLGlob
/*
PS: 搭配 gin.Context.HTML() 使用.
*/
func LoadHTMLGlob(engine *gin.Engine, pattern string) {
	engine.LoadHTMLGlob(pattern)
}

// SetHTMLTemplate
/*
PS: 搭配 gin.Context.HTML() 使用.
*/
func SetHTMLTemplate(engine *gin.Engine, templ *template.Template) {
	engine.SetHTMLTemplate(templ)
}
