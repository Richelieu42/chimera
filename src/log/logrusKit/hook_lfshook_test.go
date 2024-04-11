package logrusKit

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

// TestNewLfsHook
/*
DEBUG
*/
func TestNewLfsHook(t *testing.T) {
	logger, err := NewFileLogger("_test_lfshook.log")
	if err != nil {
		panic(err)
	}

	hook := NewLfsHook(lfshook.WriterMap{
		logrus.ErrorLevel: os.Stdout,
	}, logger.Formatter)
	logger.AddHook(hook)

	logger.Debugf("Debug %d", 0)
	logger.Infof("Info %d", 1)
	logger.Warnf("Warn %d", 2)
	logger.Errorf("Error %d", 3)
}
