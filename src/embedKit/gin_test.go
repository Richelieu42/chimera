package embedKit

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"testing"
)

//go:embed c
var efs embed.FS

func TestNewHttpFileSystemWithGin(t *testing.T) {
	engine := gin.Default()

	{
		httpFileSystem, err := NewHttpFileSystem(efs, "c")
		if err != nil {
			panic(err)
		}
		engine.StaticFS("/s", httpFileSystem)
	}

	{
		t := template.Must(template.New("").ParseFS(efs, "c/*.html"))
		engine.SetHTMLTemplate(t)

		engine.NoRoute(func(ctx *gin.Context) {
			ctx.HTML(http.StatusNotFound, "index.html", gin.H{
				"prefix": "",
				"route":  ctx.Request.URL.Path,
			})
		})
	}

	if err := engine.Run(":80"); err != nil {
		panic(engine)
	}
}
