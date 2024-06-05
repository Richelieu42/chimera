package reverseProxyKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"net/http"
	"net/http/httputil"
)

type ReverseProxy struct {
	*httputil.ReverseProxy
}

// Proxy 代理请求.
func (rp *ReverseProxy) Proxy(w http.ResponseWriter, r *http.Request) (err error) {
	if err = interfaceKit.AssertNotNil(rp, "rp"); err != nil {
		return
	}

	if obj := recover(); obj != nil {
		if err1, ok := obj.(error); ok {
			err = err1
			return
		}
		err = errorKit.Newf("recover from %v", obj)
	}
	rp.ReverseProxy.ServeHTTP(w, r)
	return
}

// wrap 不能为nil
func wrap(proxy *httputil.ReverseProxy) *ReverseProxy {
	return &ReverseProxy{
		ReverseProxy: proxy,
	}
}
