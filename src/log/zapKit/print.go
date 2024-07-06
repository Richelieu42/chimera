package zapKit

import "go.uber.org/zap"

var (
	innerL = NewLogger(nil, WithCallerSkip(1))
	innerS = innerL.Sugar()
)

func Sync() {
	_ = innerL.Sync()
	_ = innerS.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	innerL.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	innerL.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	innerL.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	innerL.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	innerL.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	innerL.DPanic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	innerL.Fatal(msg, fields...)
}
