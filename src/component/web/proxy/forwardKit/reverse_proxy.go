package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// WrapToReverseProxy *httputil.ReverseProxy => *forwardKit.ReverseProxy
func WrapToReverseProxy(reverseProxy *httputil.ReverseProxy) (*ReverseProxy, error) {
	if err := interfaceKit.AssertNotNil(reverseProxy, "reverseProxy"); err != nil {
		return nil, err
	}

	return &ReverseProxy{
		ReverseProxy: *reverseProxy,
	}, nil
}

// NewSingleHostReverseProxyWithUrl
/*
@param urlStr 		目标url
					e.g. 	"http://127.0.0.1:8000": 将请求转发给"http://127.0.0.1:8000"，请求路由不变
					e.g.1 	"http://127.0.0.1:8000/a": 将请求转发给"http://127.0.0.1:8000"，请求路由的最前面加上"/a"
@param 	errLog 		可以为nil（即无输出，但不推荐这么干）
@return !!!: Transport、ModifyResponse、ErrorHandler 等字段为nil
*/
func NewSingleHostReverseProxyWithUrl(urlStr string, errLog *log.Logger) (*ReverseProxy, error) {
	if err := strKit.AssertNotEmpty(urlStr, "urlStr"); err != nil {
		return nil, err
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, errorKit.Newf("invalid urlStr(%s)", urlStr)
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
