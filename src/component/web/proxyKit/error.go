package proxyKit

import (
	"context"
	"errors"
	"net/http"
)

// IsNegligibleError 是否是可忽略的代理error？
/*
PS: 如果返回值为true，说明请求已经结束了，无需再响应内容给http客户端了.
*/
func IsNegligibleError(err error) bool {
	if errors.Is(err, context.Canceled) {
		return true
	}
	if errors.Is(err, http.ErrAbortHandler) {
		return true
	}
	return false
}
