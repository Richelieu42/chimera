package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 定义一个处理 GET 请求的路由
	router.GET("/redirect", func(c *gin.Context) {
		// 通过 c.Redirect 方法进行重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.example.com")
	})

	// 定义一个处理根路径的路由
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// 启动 HTTP 服务，监听 8080 端口
	router.Run(":8080")
}
