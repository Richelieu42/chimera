package req111Kit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/request/httpClientKit"
)

// IsTimeoutError 是否是请求超时错误?
var IsTimeoutError func(err error) (flag bool) = httpClientKit.IsTimeoutError
