package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
)

func main() {
	r := gin.Default()

	// if Allow DirectoryIndex
	//r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
	// set prefix
	//r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))

	r.Use(static.Serve("/", static.LocalFile("/tmp", false)))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})

	ginKit.StaticDir()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
