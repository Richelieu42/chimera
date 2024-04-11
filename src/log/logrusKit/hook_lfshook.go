package logrusKit

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func NewLfsHook(output interface{}, formatter logrus.Formatter) *lfshook.LfsHook {
	return lfshook.NewHook(output, formatter)
}
