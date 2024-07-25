package forwardKit

import (
	"log"
	"net/http"
)

// ForwardToSingleHost 代理请求（请求的scheme不会变）.
/*
PS: 代理请求失败时，建议返回状态码502(http.StatusBadGateway, 网关错误).

@param errLogger	可以为nil（即无输出，但不推荐这么干）
@param url			目标url
					e.g. 	"http://127.0.0.1:8000": 将请求转发给"http://127.0.0.1:8000"，请求路由不变
					e.g.1 	"http://127.0.0.1:8000/a": 将请求转发给"http://127.0.0.1:8000"，请求路由的最前面加上"/a"
*/
func ForwardToSingleHost(w http.ResponseWriter, r *http.Request, url string, errLog *log.Logger) (err error) {
	rp, err := NewSingleHostReverseProxyWithUrl(url)
	if err != nil {
		return
	}
	rp.ErrorLog = errLog

	return ForwardByReverseProxy(w, r, rp)
}

// ForwardToHost 代理请求.
/*
PS: 代理请求失败时，建议返回状态码502(http.StatusBadGateway, 网关错误).

@param host		e.g."127.0.0.1:80"
@param errLog 	可以为nil（即无输出，但不推荐这么干）
*/
func ForwardToHost(w http.ResponseWriter, r *http.Request, host string, errLog *log.Logger, options ...DirectorOption) error {
	return ForwardToHostComplexly(w, r, host, errLog, nil, nil, options...)
}

// ForwardToHostComplexly 代理请求.
/*
PS: 代理请求失败时，建议返回状态码502(http.StatusBadGateway, 网关错误).

@param host				e.g."127.0.0.1:80"
@param errLog	 		可以为nil（即无输出，但不推荐这么干）
@param transport		可以为nil
@param modifyResponse	可以为nil
*/
func ForwardToHostComplexly(w http.ResponseWriter, r *http.Request, host string, errLog *log.Logger,
	transport http.RoundTripper, modifyResponse func(*http.Response) error, options ...DirectorOption) error {
	director, err := NewDirector(host, options...)
	if err != nil {
		return err
	}

	rp, err := NewReverseProxy(director, transport, modifyResponse, errLog, nil)
	if err != nil {
		return err
	}

	return ForwardByReverseProxy(w, r, rp)
}
