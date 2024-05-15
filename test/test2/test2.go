package main

import (
	"bytes"
	"compress/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()

	// 使用 GzipMiddleware 中间件
	r.Use(GzipMiddleware(1000)) // 设置最小长度为 1024 字节

	// 测试路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, strings.Repeat("H", 1001))
	})

	r.Run(":8080")
}

func GzipMiddleware(minLength int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Encoding", "gzip")

		// 创建一个 ResponseWriter 接口的包装器
		grw := &gzipResponseWriter{
			ResponseWriter: ctx.Writer,

			minLength: minLength,
			buffer:    bytes.NewBuffer(nil),
		}
		// 替换原始的 ResponseWriter 接口
		ctx.Writer = grw

		// 继续处理请求
		ctx.Next()

		//// 获取响应体大小
		//bodySize := grw.buffer.Len()
		//// 如果响应体大小小于配置的最小长度，则不进行压缩
		//if bodySize < minLength {
		//	grw.ResponseWriter.WriteHeader(http.StatusOK)
		//	_, _ = grw.buffer.WriteTo(ctx.Writer)
		//	return
		//}

		grw.WriteBody(ctx)
	}
}

// 自定义的 ResponseWriter 接口的包装器
type gzipResponseWriter struct {
	gin.ResponseWriter

	minLength int

	written bool
	buffer  *bytes.Buffer
}

// 重写 Write 方法，将响应体写入缓冲区
func (grw *gzipResponseWriter) Write(data []byte) (int, error) {
	grw.written = true
	return grw.buffer.Write(data)
}

// WriteBody 在请求结束时进行 gzip 压缩
func (grw *gzipResponseWriter) WriteBody(ctx *gin.Context) {
	if !grw.written {
		return
	}

	if grw.buffer.Len() < grw.minLength {
		// 不进行 gzip 压缩
		ctx.Writer.Header().Del("Content-Encoding")

		_, _ = grw.buffer.WriteTo(grw.ResponseWriter)
		return
	}

	// 进行 gzip 压缩
	writer, _ := gzip.NewWriterLevel(grw.ResponseWriter, gzip.BestSpeed)
	defer writer.Close()
	_, _ = grw.buffer.WriteTo(writer)
	//_ = writer.Flush()
}
