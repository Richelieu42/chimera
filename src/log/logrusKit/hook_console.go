package logrusKit

import (
	"github.com/sirupsen/logrus"
)

var (
	// 保护性检查
	_ logrus.Hook = (*ErrorToConsoleHook)(nil)
)

// ErrorToConsoleHook ERROR及以上的，也输出到控制台
type ErrorToConsoleHook struct {
}

func (hook *ErrorToConsoleHook) Fire(entry *logrus.Entry) error {
	// 输出到控制台
	logrus.WithFields(logrus.Fields{
		"location": GetLocation(entry.Caller),
		"level":    entry.Level.String(),
	}).Error(entry.Message)

	return nil
}

func (hook *ErrorToConsoleHook) Levels() []logrus.Level {
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
