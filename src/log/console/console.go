package console

import (
	"go.uber.org/zap"
)

var (
	innerL *zap.Logger
	innerS *zap.SugaredLogger

	l *zap.Logger
	s *zap.SugaredLogger
)

func init() {
	innerL = newLogger(1)
	innerS = innerL.Sugar()

	l = newLogger(0)
	s = l.Sugar()
}

func L() *zap.Logger {
	return l
}

func S() *zap.SugaredLogger {
	return s
}

//// ReplaceGlobalLoggers
///*
//PS: 有需要的话，应该在应用初始化时调用此方法，即在最前面设置全局logger.
//*/
//func ReplaceGlobalLoggers(logger *zap.Logger) {
//	if logger == nil {
//		return
//	}
//
//	innerL = logger
//	innerS = logger.Sugar()
//}

func Sync() {
	_ = innerL.Sync()
	_ = innerS.Sync()

	_ = l.Sync()
	_ = s.Sync()
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
