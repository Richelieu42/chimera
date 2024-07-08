package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

// NewStdLogger *zap.Logger => *log.Logger
func NewStdLogger(l *zap.Logger) *log.Logger {
	return zap.NewStdLog(l)
}

// NewStdLoggerWithLevel *zap.Logger => *log.Logger
func NewStdLoggerWithLevel(l *zap.Logger, level zapcore.Level) (*log.Logger, error) {
	return zap.NewStdLogAt(l, level)
}

func RedirectStdLog(l *zap.Logger) func() {
	return zap.RedirectStdLog(l)
}

func RedirectStdLogWithLevel(l *zap.Logger, level zapcore.Level) (func(), error) {
	return zap.RedirectStdLogAt(l, level)
}
