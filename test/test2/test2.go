package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.MustSetUp(nil)

	// 示例日志输出
	logrus.WithFields(logrus.Fields{
		"a": 0,
		"b": 1,
	}).Info("This is an info message1111111111111")
	logrus.Warn("This is a warning message")
	logrus.Error("This is an error message")
}
