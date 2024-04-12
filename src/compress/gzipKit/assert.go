package gzipKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/funcKit"
)

func AssertValidLevel(level int) error {
	if !IsValidLevel(level) {
		return errorKit.NewfWithSkip(1, "[%s] invalid gzip compression level(%d)", funcKit.GetFuncName(1), level)
	}
	return nil
}
