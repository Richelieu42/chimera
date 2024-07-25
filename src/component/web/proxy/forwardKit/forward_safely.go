package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"net/http"
	"net/http/httputil"
)

// ForwardByReverseProxy !!!: 不要直接使用 httputil.ReverseProxy 的 ServeHTTP 方法.
func ForwardByReverseProxy(reverseProxy *httputil.ReverseProxy, w http.ResponseWriter, r *http.Request) (err error) {
	if err = interfaceKit.AssertNotNil(reverseProxy, "reverseProxy"); err != nil {
		return
	}
	rp := *reverseProxy

	/*
		设置 ErrorHandler 字段，以免: 代理请求失败时，返回的err == nil.
		PS: 并不会影响外部的rp，因为不是"指针类型".
	*/
	old := rp.ErrorHandler
	rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, e error) {
		err = e
		if rp.ErrorLog != nil {
			rp.ErrorLog.Printf("Fail to forward request, error: %s", err.Error())
		}

		if old != nil {
			old(w, r, e)
		}
	}

	// Richelieu: try to reset http.Request.Body
	if err = httpKit.TryToResetRequestBody(r); err != nil {
		return
	}

	// Richelieu: 请求转发前再检查下，以防请求已经被取消了
	if err = r.Context().Err(); err != nil {
		return
	}

	// 真正转发请求
	rp.ServeHTTP(w, r)
	return
}
