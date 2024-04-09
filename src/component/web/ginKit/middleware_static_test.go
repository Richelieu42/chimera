package ginKit

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestNewStaticMiddleware(t *testing.T) {
	engine := gin.Default()

	engine.Use(NewStaticMiddleware("/s", "_test/", true))

	engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "OK")
	})

	// Listen and Server in 0.0.0.0:80
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
