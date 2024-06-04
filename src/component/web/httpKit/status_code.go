package httpKit

import (
	"net/http"
)

var (
	// GetStatusText 获取 http状态码 的描述文本.
	GetStatusText func(code int) string = http.StatusText
)

// IsValidStatusCode 响应的 http状态码 是否有效？
/*
PS:
(1) 判断参考了 jQuery;
(2) 正常情况下，响应的http状态码有效的情况下，才会去读取响应的body.
*/
func IsValidStatusCode(code int) bool {
	return code >= 200 && code < 300 || code == http.StatusNotModified
}
