package main

import (
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

type hook struct {
	prefix string
}

func (hook *hook) Fire(entry *logrus.Entry) error {
	entry.Message += "==="

	return nil
}

func (hook *hook) Levels() []logrus.Level {
	// 只有 INFO、WARN 级别
	return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel}

	//return logrus.AllLevels
}

func main() {
	logger, err := logrusKit.NewFileLogger("_test.log")
	if err != nil {
		panic(err)
	}

	logger.AddHook(&hook{})

	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}
