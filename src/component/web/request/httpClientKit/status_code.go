package httpClientKit

import (
	"net/http"
)

// IsCodeValid 响应的http状态码 是否有效？
/*
PS:
(1) 判断参考了jQuery;
(2) 正常情况下，响应的http状态码有效的情况下，才会去读取响应的body.
*/
func IsCodeValid(code int) bool {
	return code >= 200 && code < 300 || code == http.StatusNotModified
}
