package proxyKit

import (
	"context"
	"errors"
	"net"
	"net/http"
)

// IsNegligibleError 是否是可忽略的代理error？
/*
Deprecated: Use forwardKit instead.

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

// IsProxyDialError 代理请求返回的error，是否是因为dial目标地址失败？
/*
Deprecated: Use forwardKit instead.
*/
func IsProxyDialError(err error) bool {
	opErr := &net.OpError{}
	if errors.As(err, &opErr) {
		return opErr.Op == "dial"
	}
	return false
}
