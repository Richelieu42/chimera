package embedKit

import (
	"embed"
	"github.com/gin-gonic/gin"
	"testing"
)

//go:embed c
var efs embed.FS

func TestNewHttpFileSystemWithGin(t *testing.T) {
	engine := gin.Default()

	if err := engine.Run(":80"); err != nil {
		panic(engine)
	}
}
