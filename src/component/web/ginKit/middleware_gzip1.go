package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"
)

// NewGzipMiddleware1
/*
PS: 涉及多个服务（请求转发）的场景下，	(1) 最外层的务使用gzip压缩;
									(2) 内层的服务不使用gzip压缩.

@param level			压缩级别
@param minContentLength	(1) 触发gzip的最小内容长度
						(2) 单位: byte
						(3) 必须 > 0
*/
func NewGzipMiddleware1(level int, minContentLength int64) gin.HandlerFunc {
	gzipHandler := gzip.NewHandler(gzip.Config{
		CompressionLevel: level,
		MinContentLength: minContentLength,
		RequestFilter: []gzip.RequestFilter{
			gzip.NewCommonRequestFilter(),
			//gzip.DefaultExtensionFilter(),
		},
		ResponseHeaderFilter: []gzip.ResponseHeaderFilter{
			gzip.NewSkipCompressedFilter(),
			gzip.DefaultContentTypeFilter(),
		},
	})
	return gzipHandler.Gin
}
