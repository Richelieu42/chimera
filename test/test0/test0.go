package main

import (
	"github.com/imroc/req/v3"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
)

func main() {
	req.DefaultClient()
	req.SetDefaultClient()

	req.C().ImpersonateChrome()
	req.C().ImpersonateFirefox()
	req.C().ImpersonateSafari()
	req.C().Impersona

	req.DevMode()
	req.EnableForceHTTP1()

	req.AddCommonRetryCondition(func(resp *req.Response, err error) bool {

	})
}
