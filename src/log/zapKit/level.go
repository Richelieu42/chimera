package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"go.uber.org/zap"
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

// CanPrintSpecifiedLevel logger是否会打印 指定级别lv 的日志？
func CanPrintSpecifiedLevel(logger *zap.Logger, lv zapcore.Level) bool {
	if logger == nil {
		return false
	}

	return lv >= logger.Level()
}
