package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"net/http"
	"net/url"
)

func NewSingleHostDirector(u *url.URL) (func(r *http.Request), error) {
	if err := interfaceKit.AssertNotNil(u, "u"); err != nil {
		return nil, err
	}

	rp, err := NewSingleHostReverseProxy(u)
	if err != nil {
		return nil, err
	}
	return rp.Director, nil
}
