package httpKit

import "net/http"

// GetOrigin 获取请求的Origin
func GetOrigin(header http.Header) string {
	return GetHeader(header, HeaderOrigin)
}
