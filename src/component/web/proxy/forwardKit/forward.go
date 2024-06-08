package forwardKit

import (
	"log"
	"net/http"
)

// ForwardToSingleHost 代理请求.
/*
@param errLogger	可以为nil（即无输出，但不推荐这么干）
@param url			目标url
					e.g. 	"http://127.0.0.1:8000": 将请求转发给"http://127.0.0.1:8000"，请求路由不变
					e.g.1 	"http://127.0.0.1:8000/a": 将请求转发给"http://127.0.0.1:8000"，请求路由的最前面加上"/a"
*/
func ForwardToSingleHost(w http.ResponseWriter, r *http.Request, url string, errLog *log.Logger) (err error) {
	rp, err := NewSingleHostReverseProxyWithUrl(url, errLog)
	if err != nil {
		return
	}
	return rp.Forward(w, r)
}

// ForwardToHost
/*
@param errorLog 可以为nil（即无输出，但不推荐这么干）
*/
func ForwardToHost(w http.ResponseWriter, r *http.Request, host string, errLog *log.Logger, options ...DirectorOption) error {
	return ForwardToHostComplexly(w, r, host, errLog, nil, nil, options...)
}

// ForwardToHostComplexly
/*
@param errLogger 		可以为nil（即无输出，但不推荐这么干）
@param transport		可以为nil
@param modifyResponse	可以为nil
*/
func ForwardToHostComplexly(w http.ResponseWriter, r *http.Request, host string, errLog *log.Logger, transport http.RoundTripper, modifyResponse func(*http.Response) error, options ...DirectorOption) error {
	director, err := NewDirector(host, options...)
	if err != nil {
		return err
	}
	rp, err := NewReverseProxy(director, transport, modifyResponse, errLog, nil)
	if err != nil {
		return err
	}
	return rp.Forward(w, r)
}
