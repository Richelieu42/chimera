package reverseProxyKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

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
