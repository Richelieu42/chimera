package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestNewCore(t *testing.T) {
	// 创建错误日志级别的核心
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})

	NewCore(nil, nil, errorLevel)
}
