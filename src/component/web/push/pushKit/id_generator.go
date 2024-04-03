package pushKit

import "github.com/richelieu-yang/chimera/v3/src/idKit"

func DefaultIdGenerator() func() (string, error) {
	return func() (string, error) {
		return idKit.NewXid(), nil
	}
}
