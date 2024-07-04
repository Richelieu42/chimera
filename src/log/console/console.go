package console

import (
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
)

var (
	l *zap.Logger
	s *zap.SugaredLogger
)

func init() {
	/*
		WithCallerSkip(1): 跳过1层，因为进行了1层封装
	*/
	l = zapKit.NewLogger(nil, zapKit.WithCallerSkip(1))
	s = l.Sugar()
}

func Sync() {
	_ = l.Sync()
	_ = s.Sync()
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

// Debugf 格式化输出的信息日志，类似于fmt.Printf，可以使用格式化字符串.
func Debugf(template string, args ...interface{}) {
	s.Debugf(template, args...)
}

// Debugw 结构化输出的信息日志，使用键值对的方式输出，更加适合记录结构化数据.
/*
@param keysAndValues e.g. "key", "value", "flag", true
*/
func Debugw(msg string, keysAndValues ...interface{}) {
	s.Debugw(msg, keysAndValues...)
}

// Debugln
/*
PS: Spaces are always added between arguments.（传参间会加上" "）
*/
func Debugln(args ...interface{}) {
	s.Debugln(args...)
}

func Infof(template string, args ...interface{}) {
	s.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	s.Infow(msg, keysAndValues...)
}

func Infoln(args ...interface{}) {
	s.Infoln(args...)
}

func Warnf(template string, args ...interface{}) {
	s.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	s.Warnw(msg, keysAndValues...)
}

func Warnln(args ...interface{}) {
	s.Warnln(args...)
}

func Errorf(template string, args ...interface{}) {
	s.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	s.Errorw(msg, keysAndValues...)
}

func Errorln(args ...interface{}) {
	s.Errorln(args...)
}

func DPanicf(template string, args ...interface{}) {
	s.DPanicf(template, args...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	s.DPanicw(msg, keysAndValues...)
}

func DPanicln(args ...interface{}) {
	s.DPanicln(args...)
}

func Panicf(template string, args ...interface{}) {
	s.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	s.Panicw(msg, keysAndValues...)
}

func Panicln(args ...interface{}) {
	s.Panicln(args...)
}

func Fatalf(template string, args ...interface{}) {
	s.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	s.Fatalw(msg, keysAndValues...)
}

func Fatalln(args ...interface{}) {
	s.Fatalln(args...)
}
