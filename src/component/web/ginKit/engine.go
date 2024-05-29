package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
)

func NewEngine() *gin.Engine {
	engine := gin.New()

	// true: enable fallback Context.Deadline(), Context.Done(), Context.Err() and Context.Value() when Context.Request.Context() is not nil
	engine.ContextWithFallback = true

	// 默认: true
	engine.ForwardedByClientIP = true
	// 默认: []string{"X-Forwarded-For", "X-Real-IP"}
	engine.RemoteIPHeaders = httpKit.RemoteIPHeaders

	// 默认: true
	engine.RedirectTrailingSlash = true

	/*
		MaxMultipartMemory只是限制内存，不是针对文件上传文件大小，即使文件大小比这个大，也会写入临时文件。
		默认32MiB，并不涉及"限制上传文件的大小"，原因：上传的文件s按顺序存入内存中，累加大小不得超出 32Mb ，最后累加超出的文件就存入系统的临时文件中。非文件字段部分不计入累加。所以这种情况，文件上传是没有任何限制的。
		参考: https://studygolang.com/articles/22643
	*/
	engine.MaxMultipartMemory = 64 << 20

	return engine
}
