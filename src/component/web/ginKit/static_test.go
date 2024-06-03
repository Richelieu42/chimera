package ginKit

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

//go:embed _test
var tmpFs embed.FS

/*
url: http://127.0.0.1/a/b/_test/a.txt
*/
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

/*
url: http://127.0.0.1/a/b/a.txt
*/
func TestStaticDir(t *testing.T) {
	engine := gin.Default()

	if err := StaticDir(engine, "/a/b", "_test", false); err != nil {
		panic(err)
	}

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
