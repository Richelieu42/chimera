package ginKit

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

//go:embed _test
var tmpFs embed.FS

func TestStaticFS(t *testing.T) {
	engine := gin.Default()

	group := engine.Group("a")
	if err := StaticFS(group, "b", http.FS(tmpFs)); err != nil {
		panic(err)
	}

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
