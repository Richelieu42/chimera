package oxyKit

import (
	"github.com/sirupsen/logrus"
	"github.com/vulcand/oxy/v2/utils"
)

type loggerWrapper struct {
	utils.Logger
	logger *logrus.Logger
}

func NewLogger(logrusLogger *logrus.Logger) (logger utils.Logger) {
	if logrusLogger == nil {
		logger = &utils.NoopLogger{}
		return
	}

	return
}
