package netKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
)

// ProcessAddresses
/*
@return 如果第二个返回值为nil，那么 len(第一个返回值) > 0
*/
func ProcessAddresses(addresses []string) ([]string, error) {
	addresses = sliceKit.PolyfillStringSlice(addresses)
	if err := sliceKit.AssertNotEmpty(addresses, "addresses"); err != nil {
		return nil, err
	}

	s := make([]string, 0, len(addresses))
	for _, address := range addresses {
		a, err := ParseToAddress(address)
		if err != nil {
			return nil, err
		}
		s = append(s, a.String())
	}
	return s, nil
}
