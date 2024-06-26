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

func Info(args ...interface{}) {
	sl.Info(args...)
}

func Infof(template string, args ...interface{}) {
	sl.Infof(template, args...)
}
