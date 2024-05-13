package proxyKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"net/http"
)

func Mark(req *http.Request) {
	httpKit.SetHeader(req.Header, httpKit.HeaderChimeraCors, "1")
}

func IsMarked(req *http.Request) bool {
	return httpKit.GetHeader(req.Header, httpKit.HeaderChimeraCors) == "1"
}
