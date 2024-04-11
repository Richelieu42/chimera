package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"

	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

type hook struct {
	prefix string
}

func (hook *hook) Fire(entry *logrus.Entry) error {
	logrus.Info(entry.Message)

	return nil
}

func (hook *hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		//logrus.WarnLevel,
		//logrus.InfoLevel,
		//logrus.DebugLevel,
		//logrus.TraceLevel,
	}

	//return logrus.AllLevels
}

func main() {
	logger, err := logrusKit.NewFileLogger("_test.log")
	if err != nil {
		panic(err)
	}

	logger.AddHook(&hook{})

	logger.Debugf("Debug %d", 0)
	logger.Infof("Info %d", 1)
	logger.Warnf("Warn %d", 2)
	logger.Errorf("Error %d", 3)
	logger.Panicf("Panic %d", 4)
}
