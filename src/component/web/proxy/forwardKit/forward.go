package forwardKit

import (
	"log"
	"net/http"
)

// ForwardToUrl 代理请求到目标url.
/*
@param errLogger	可以为nil（即无输出，但不推荐这么干）
@param url			目标url
*/
func ForwardToUrl(w http.ResponseWriter, r *http.Request, url string, errLog *log.Logger) (err error) {
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
