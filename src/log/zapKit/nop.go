package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewNopCore() zapcore.Core {
	return zapcore.NewNopCore()
}

// NewNopLogger 不进行实际日志记录操作.
func NewNopLogger() *zap.Logger {
	return zap.NewNop()
}

// NewNopSugaredLogger 不进行实际日志记录操作.
func NewNopSugaredLogger() *zap.SugaredLogger {
	return zap.NewNop().Sugar()
}
