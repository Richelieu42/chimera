package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/urlKit"
	"log"
	"net/http"
)

// ForwardToUrl 代理请求到目标url.
/*
@param errorLog 可以为nil，即无输出
@param url		目标url
*/
func ForwardToUrl(w http.ResponseWriter, r *http.Request, errorLog *log.Logger, url string) (err error) {
	u, err := urlKit.Parse(url)
	if err != nil {
		err = errorKit.Wrapf(err, "invalid url(%s)", url)
		return
	}

	rp, err := NewSingleHostReverseProxy(u)
	if err != nil {
		return
	}
	rp.ErrorLog = errorLog
	rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err1 error) {
		err = err1
		if errorLog != nil {
			errorLog.Printf("Fail to forward request to url(%s), error: %s.", url, err.Error())
		}
	}

	err = rp.Forward(w, r)
	return
}

func ForwardToHost(w http.ResponseWriter, r *http.Request, errorLog *log.Logger, host string, options ...DirectorOption) (err error) {
	director, err := NewDirector(host, options...)
	if err != nil {
		return
	}
	rp, err := NewReverseProxy(director, nil, nil, errorLog, nil)
	if err != nil {
		return
	}
	rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err1 error) {
		err = err1
		if errorLog != nil {
			errorLog.Printf("Fail to forward request to host(%s), error: %s.", host, err.Error())
		}
	}

	err = rp.Forward(w, r)
	return
}
