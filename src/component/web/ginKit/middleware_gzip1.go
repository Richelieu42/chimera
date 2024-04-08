package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmu42/gzip"
)

// NewGzipMiddleware1
/*
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
