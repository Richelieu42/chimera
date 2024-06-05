package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxy struct {
	*httputil.ReverseProxy
}

// Forward 代理请求.
func (rp *ReverseProxy) Forward(w http.ResponseWriter, r *http.Request) (err error) {
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

	/* Richelieu: try to reset http.Request.Body */
	if err = httpKit.TryToResetRequestBody(r); err != nil {
		return
	}
	/* Richelieu: 请求转发前再检查下，以防请求已经被取消了 */
	if err = r.Context().Err(); err != nil {
		return err
	}

	// 真正的请求转发
	rp.ReverseProxy.ServeHTTP(w, r)
	return
}

// wrap 不能为nil
func wrap(proxy *httputil.ReverseProxy) *ReverseProxy {
	return &ReverseProxy{
		ReverseProxy: proxy,
	}
}

// NewSingleHostReverseProxy
/*
	@param target 不能为nil，否则会panic
*/
func NewSingleHostReverseProxy(target *url.URL) (*ReverseProxy, error) {
	if err := interfaceKit.AssertNotNil(target, "target"); err != nil {
		return nil, err
	}

	tmp := httputil.NewSingleHostReverseProxy(target)
	return wrap(tmp), nil
}

// NewReverseProxy
/*
PS: 对于 httputil.ReverseProxy 结构体，Rewrite 和 Director 只能有一个非nil.

@param director			不能为nil!!!
@param transport		可以为nil
@param modifyResponse	可以为nil
@param errLog			可以为nil
@param errHandler		可以为nil
*/
func NewReverseProxy(director func(*http.Request), transport http.RoundTripper, modifyResponse func(*http.Response) error, errLog *log.Logger, errHandler func(http.ResponseWriter, *http.Request, error)) (*ReverseProxy, error) {
	if err := interfaceKit.AssertNotNil(director, "director"); err != nil {
		return nil, err
	}

	tmp := &httputil.ReverseProxy{
		Director:       director,
		Transport:      transport,
		ModifyResponse: modifyResponse,
		//BufferPool:     nil,
		//FlushInterval:  0,
		ErrorLog:     errLog,
		ErrorHandler: errHandler,
	}
	return wrap(tmp), nil
}
