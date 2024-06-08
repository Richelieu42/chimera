package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxy struct {
	*httputil.ReverseProxy
}

// Forward 请求转发（代理请求; proxy）.
/*
@return !!!: 此方法的返回值需要注意，就算为nil，也有可能请求转发失败，还需要看 http.ReverseProxy.ErrorHandler.
*/
func (rp *ReverseProxy) Forward(w http.ResponseWriter, r *http.Request) (err error) {
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

	// 设置ErrorHandler字段，以免请求转发失败时没有error
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

// WrapToReverseProxy *httputil.ReverseProxy => *forwardKit.ReverseProxy
func WrapToReverseProxy(reverseProxy *httputil.ReverseProxy) (*ReverseProxy, error) {
	if err := interfaceKit.AssertNotNil(reverseProxy, "reverseProxy"); err != nil {
		return nil, err
	}

	return &ReverseProxy{
		ReverseProxy: reverseProxy,
	}, nil
}

// NewSingleHostReverseProxyWithUrl
/*
@param targetUrl 	e.g."http://127.0.0.1:8000/test"
@param 	errLog 		可以为nil（即无输出，但不推荐这么干）
@return !!!: Transport、ModifyResponse、ErrorHandler 等字段为nil
*/
func NewSingleHostReverseProxyWithUrl(targetUrl string, errLog *log.Logger) (*ReverseProxy, error) {
	if err := strKit.AssertNotEmpty(targetUrl, "targetUrl"); err != nil {
		return nil, err
	}
	u, err := url.Parse(targetUrl)
	if err != nil {
		return nil, errorKit.Newf("invalid targetUrl(%s)", targetUrl)
	}

	return NewSingleHostReverseProxy(u, errLog)
}

// NewSingleHostReverseProxy
/*
@param target 不能为nil，否则会panic
@param errLog 可以为nil（即无输出，但不推荐这么干）
@return !!!: Transport、ModifyResponse、ErrorHandler 等字段为nil
*/
func NewSingleHostReverseProxy(u *url.URL, errLog *log.Logger) (*ReverseProxy, error) {
	if err := interfaceKit.AssertNotNil(u, "u"); err != nil {
		return nil, err
	}

	tmp := httputil.NewSingleHostReverseProxy(u)
	tmp.ErrorLog = errLog

	return WrapToReverseProxy(tmp)
}

// NewReverseProxy
/*
PS: 对于 httputil.ReverseProxy 结构体，Rewrite 和 Director 只能有一个非nil.

@param director			不能为nil!!!
@param transport		可以为nil
@param modifyResponse	可以为nil
@param errLog			可以为nil（即无输出，但不推荐这么干）
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
	return WrapToReverseProxy(tmp)
}
