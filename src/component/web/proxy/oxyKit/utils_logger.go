package oxyKit

import (
	"github.com/sirupsen/logrus"
	"github.com/vulcand/oxy/v2/utils"
)

type loggerWrapper struct {
	*logrus.Logger
}

func (l loggerWrapper) Debug(msg string, args ...any) {
	l.Debugf(msg, args...)
}

func (l loggerWrapper) Info(msg string, args ...any) {
	l.Infof(msg, args...)
}

func (l loggerWrapper) Warn(msg string, args ...any) {
	l.Warnf(msg, args...)
}

func (l loggerWrapper) Error(msg string, args ...any) {
	l.Errorf(msg, args...)
}

func NewLogger(logrusLogger *logrus.Logger) (logger utils.Logger) {
	if logrusLogger == nil {
		// 不输出
		return &utils.NoopLogger{}
	}
	return &loggerWrapper{logrusLogger}
}
