package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"go.uber.org/zap/zapcore"
	"strings"
)

// StringToLevel
/*
PS: 默认 DEBUG 级别.
*/
func StringToLevel(str string) (zapcore.Level, error) {
	if strKit.IsBlank(str) {
		return zapcore.DebugLevel, nil
	}

	str = strings.ToLower(str)
	return zapcore.ParseLevel(str)
}
