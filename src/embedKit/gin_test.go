package embedKit

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
	"testing"
)

//go:embed c
var efs embed.FS

func TestNewHttpFileSystemWithGin(t *testing.T) {
	engine := gin.Default()

	ginKit.StaticFile()

	if err := engine.Run(":80"); err != nil {
		panic(engine)
	}
}
