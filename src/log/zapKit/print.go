package zapKit

import "go.uber.org/zap"

var (
	l = NewLogger(nil, WithCallerSkip(0))
	s = l.Sugar()

	innerL = NewLogger(nil, WithCallerSkip(1))
	innerS = innerL.Sugar()
)

func L() *zap.Logger {
	return l
}

func S() *zap.SugaredLogger {
	return s
}

func Sync() {
	_ = l.Sync()
	_ = s.Sync()

	_ = innerL.Sync()
	_ = innerS.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	innerL.Debug(msg, fields...)
}

// Info
/*
@param fields 输出循序与传参顺序一致（并不会按字母排序）
*/
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
