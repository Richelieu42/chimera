package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"net/http"
	"net/http/httputil"
)

type ReverseProxy struct {
	// !!!: 此处不能是 *httputil.ReverseProxy，因为会在 Forward 方法体内修改receiver的字段，但不希望修改方法体外的.
	httputil.ReverseProxy
}

// Forward 请求转发（代理请求; proxy）.
/*
!!!: 此方法的receiver不能为指针类型，因为会在方法体内修改receiver的字段，但不希望修改方法体外的.
*/
func (rp ReverseProxy) Forward(w http.ResponseWriter, r *http.Request) (err error) {
	if err = interfaceKit.AssertNotNil(rp, "rp"); err != nil {
		return
	}

	// 主要针对: http.ReverseProxy.ServeHTTP() 中的 panic(http.ErrAbortHandler)
	if obj := recover(); obj != nil {
		if err1, ok := obj.(error); ok {
			err = err1
			return
		}
		err = errorKit.Newf("recover from %v", obj)
	}

	// ### 设置ErrorHandler字段，以免请求转发失败时 err == nil，但并不会影响外部的rp，因为receiver的类型不是指针类型
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

	// 真正的请求转发
	rp.ReverseProxy.ServeHTTP(w, r)
	return
}
