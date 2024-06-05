package forwardKit

import (
	"context"
	"errors"
	"net"
	"net/http"
)

// IsInterruptedError 是否是被中断error？（e.g. 请求被取消...）
/*
PS: 如果返回值为true，说明请求已经结束了，无需再响应内容给http客户端了.
*/
func IsInterruptedError(err error) bool {
	if errors.Is(err, context.Canceled) {
		return true
	}
	if errors.Is(err, http.ErrAbortHandler) {
		return true
	}
	return false
}

// IsProxyDialError 代理请求返回的error，是否是因为dial目标地址失败？
func IsProxyDialError(err error) bool {
	opErr := &net.OpError{}
	if errors.As(err, &opErr) {
		return opErr.Op == "dial"
	}
	return false
}
