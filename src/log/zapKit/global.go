package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"go.uber.org/zap"
)

var (
	globalMutex = new(mutexKit.RWMutex)

	l  *zap.Logger
	sl *zap.SugaredLogger
)

func init() {
	/*
		WithCallerSkip(1): 跳过1层，因为进行了1层封装
	*/
	l = NewLogger(WithCallerSkip(1))
	sl = l.Sugar()
}

func Sync() {
	_ = l.Sync()
	_ = sl.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	l.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	l.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	l.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	l.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	l.DPanic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	l.Fatal(msg, fields...)
}
