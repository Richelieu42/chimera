package zapKit

import "go.uber.org/zap"

var (
	sl *zap.SugaredLogger
)

func init() {
	/*
		WithCallerSkip(1): 跳过1层，因为进行了1层封装
	*/
	sl = NewSugarLogger(WithCallerSkip(1))
}

func Sync() error {
	return sl.Sync()
}

func Info(args ...interface{}) {
	sl.Info(args...)
}

// Infof 格式化输出的信息日志，类似于 fmt.Printf
func Infof(template string, args ...interface{}) {
	sl.Infof(template, args...)
}

// Infow 结构化输出的信息日志
func Infow(msg string, keysAndValues ...interface{}) {
	sl.Infow(msg, keysAndValues...)
}

func Infoln(args ...interface{}) {
	sl.Infoln(args...)
}
