package logrusKit

import (
	"github.com/sirupsen/logrus"
)

var (
	// 保护性检查
	_ logrus.Hook = (*prefixHook)(nil)
)

type prefixHook struct {
	prefix string
}

func (hook *prefixHook) Fire(entry *logrus.Entry) error {
	entry.Message = hook.prefix + entry.Message

	return nil
}

func (hook *prefixHook) Levels() []logrus.Level {
	// 只有 INFO、WARN 级别，才会触发 Fire()
	//return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel}

	// 所有级别，都会触发 Fire()
	return logrus.AllLevels
}
