package proxyKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"net/http"
)

// Mark 在请求头加个标记，以证明此请求被 chimera 代理过，可能被业务处理过（e.g. 允许跨域）.
/*
@param header 请求头
*/
func Mark(header http.Header) {
	httpKit.SetHeader(header, httpKit.HeaderChimeraProxy, "c")
}

// IsMarked 此请求是否被 chimera 代理过？
/*
@param header 请求头
*/
func IsMarked(header http.Header) bool {
	return httpKit.GetHeader(header, httpKit.HeaderChimeraProxy) == "c"
}
