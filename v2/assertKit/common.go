package assertKit

import (
	"github.com/richelieu42/chimera/v2/core/errorKit"
)

func NotNil(obj interface{}, variableName string) error {
	if obj == nil {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable(name: %s) is nil", variableName)
	}
	return nil
}

func Nil(obj interface{}, variableName string) error {
	if obj != nil {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable(name: %s) isn't nil", variableName)
	}
	return nil
}
