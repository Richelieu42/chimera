package sliceKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

func AssertNotEmpty[T any](s []T) error {
	if len(s) == 0 {
		if s == nil {
			return errorKit.Simple("[%s] s == nil", funcKit.GetFuncName(1))
		}
		return errorKit.Simple("[%s] len(s) == 0", funcKit.GetFuncName(1))
	}
	return nil
}
