package logrusKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/sirupsen/logrus"
)

//// SetConsoleLevel 设置控制台输出的日志级别
//func SetConsoleLevel(level logrus.Level) {
//	logrus.SetLevel(level)
//}

//func SetLevel(logger *logrus.Logger, level logrus.Level) {
//	if logger != nil {
//		logger.SetLevel(level)
//	}
//}

// ParseLevel string => logrus.Level
/*
PS: 默认日志级别 Debug.
*/
func ParseLevel(str string) (logrus.Level, error) {
	if strKit.IsBlank(str) {
		return logrus.DebugLevel, nil
	}
	return logrus.ParseLevel(str)
}
