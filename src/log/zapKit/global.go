package zapKit

import "go.uber.org/zap"

var (
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

func Sync() error {
	return sl.Sync()
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

//// Infof 格式化输出的信息日志，类似于 fmt.Printf
//func Infof(template string, args ...interface{}) {
//	sl.Infof(template, args...)
//}
//
//// Infow 结构化输出的信息日志
//func Infow(msg string, keysAndValues ...interface{}) {
//	sl.Infow(msg, keysAndValues...)
//}
//
//func Infoln(args ...interface{}) {
//	sl.Infoln(args...)
//}
