package ginKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewOptionsMiddleware
/*
Deprecated: NewCorsMiddleware 已经自带对 OPTIONS请求 的处理了.

PS:
(1) 需要先使用 NewCorsMiddleware;
(2) 就算 NoRoute || NoMethod，也会走到中间件.
*/
func NewOptionsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {
			// 预检请求的返回结果能缓存多久？24h
			ctx.Header("Access-Control-Max-Age", "86400")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Header("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
			ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

			ctx.Header("Content-Type", "text/plain; charset=utf-8")
			//ctx.Header("Content-Length", "0")

			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}
