package logrusKit

import (
	"github.com/sirupsen/logrus"
)

var (
	// 保护性检查
	_ logrus.Hook = (*defaultPrefixHook)(nil)
)

type defaultPrefixHook struct {
	prefix string
}

func (hook *defaultPrefixHook) Fire(entry *logrus.Entry) error {
	entry.Message = hook.prefix + entry.Message

	return nil
}

func (hook *defaultPrefixHook) Levels() []logrus.Level {
	//// 只有 INFO、WARN 级别
	//return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel}

	return logrus.AllLevels
}
