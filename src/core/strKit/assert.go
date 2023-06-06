package strKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

func AssertNotEmpty(str string, name string) error {
	if IsEmpty(str) {
		return errorKit.NewSkipf(1, "[%s] param(name: %s, type: %s) is empty",
			funcKit.GetFuncName(1), name, "string")
	}
	return nil
}

func AssertNotBlank(str string, name string) error {
	if IsBlank(str) {
		return errorKit.NewSkipf(1, "[%s] param(name: %s, type: %s) is blank",
			funcKit.GetFuncName(1), name, "string")
	}
	return nil
}
