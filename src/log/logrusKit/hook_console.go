package logrusKit

import (
	"github.com/sirupsen/logrus"
)

var (
	// 保护性检查
	_ logrus.Hook = (*errorToConsoleHook)(nil)
)

func NewErrorToConsoleHook() logrus.Hook {
	return &errorToConsoleHook{}
}

// errorToConsoleHook 效果: ERROR及以上的，也输出到控制台
type errorToConsoleHook struct {
}

func (hook *errorToConsoleHook) Fire(entry *logrus.Entry) error {
	// 输出到控制台
	logrus.WithFields(logrus.Fields{
		"entryLocation": GetLocation(entry.Caller),
		"entryLevel":    entry.Level.String(),
	}).Error(entry.Message)

	return nil
}

func (hook *errorToConsoleHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		//logrus.WarnLevel,
		//logrus.InfoLevel,
		//logrus.DebugLevel,
		//logrus.TraceLevel,
	}
}
