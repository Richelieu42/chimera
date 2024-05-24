package configKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
)

// PolyfillHosts
/*
PS:
(1) host包含 ip 和 port. e.g."127.0.0.1:80"

@param minCount 至少要有几个元素？

@return 第1个返回值: 可能是一个新的slice实例
*/
func PolyfillHosts(hosts []string, minCount int) ([]string, error) {
	if minCount <= 0 {
		minCount = 1
	}

	hosts = sliceKit.PolyfillStringSlice(hosts)
	if err := sliceKit.AssertNotEmpty(hosts, "hosts"); err != nil {
		return nil, err
	}
	tag := fmt.Sprintf("required,gte=%d,unique,dive,hostname_port", minCount)
	if err := validateKit.Var(hosts, tag); err != nil {
		err = errorKit.Wrapf(err, "hosts is invalid")
		return nil, err
	}
	return hosts, nil
}
