package ginKit

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// NewGzipMiddleware
/*
PS: 涉及多个服务（请求转发）的场景下，(1) 最外层的务使用gzip压缩;
								(2) 内层的服务不使用gzip压缩.
*/
var NewGzipMiddleware func(level int, options ...gzip.Option) gin.HandlerFunc = gzip.Gzip
