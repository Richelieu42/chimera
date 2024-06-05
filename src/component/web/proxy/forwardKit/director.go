package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/urlKit"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
	"net/http"
)

// NewDirector
/*
@param targetHost hostname || hostname:port
*/
func NewDirector(targetHost string, options ...DirectorOption) (dewDirector func(req *http.Request), err error) {
	if err = validateKit.Var(targetHost, "hostname|ipv4|hostname_port"); err != nil {
		err = errorKit.Newf("invalid targetHost(%s)", targetHost)
		return
	}

	opts := loadOptions(options...)

	dewDirector = func(req *http.Request) {
		req.URL.Scheme = opts.scheme
		req.URL.Host = targetHost
		if opts.requestUrlPath != nil {
			req.URL.Path = *opts.requestUrlPath
		}

		// 可能会修改 r.URL.RawQuery
		if opts.overrideQueryParams != nil {
			urlKit.OverrideRawQuery(req.URL, opts.overrideQueryParams)
		} else if opts.extraQueryParams != nil {
			urlKit.AddToRawQuery(req.URL, opts.extraQueryParams)
		}
	}
	return
}
