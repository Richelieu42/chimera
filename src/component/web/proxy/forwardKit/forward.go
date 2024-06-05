package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/urlKit"
	"log"
	"net/http"
)

// ForwardToUrl 代理请求到指定url.
/*
@param errorLog 可以为nil，即无输出
*/
func ForwardToUrl(w http.ResponseWriter, r *http.Request, url string, errorLog *log.Logger) (err error) {
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

	}
	err = rp.Forward(w, r)
	return
}
