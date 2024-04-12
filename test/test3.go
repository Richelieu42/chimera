package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 创建Gzip实例并设置阈值（例如，仅对大于1KB的响应进行压缩）
	gz := gzip.Gzip(gzip.BestSpeed)
	gz.MinLength = 1024 // 设置最小压缩长度阈值为1KB

	// 将Gzip中间件添加到Gin引擎中
	r.Use(gz)

	// ... 其他路由和中间件配置 ...

	r.Run(":8080")
}
