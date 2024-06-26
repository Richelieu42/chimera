package zapKit

import "go.uber.org/zap"

var (
	sl *zap.SugaredLogger
)

func init() {
	sl = NewSugarLogger()
}

func Info(args ...interface{}) {
	sl.Info(args...)
}

func Infof(template string, args ...interface{}) {
	sl.Infof(template, args...)
}
