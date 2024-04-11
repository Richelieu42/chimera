package logrusKit

import (
	"testing"
)

func TestNewErrorToConsoleHook(t *testing.T) {
	logger, err := NewFileLogger("_test.log")
	if err != nil {
		panic(err)
	}

	logger.AddHook(NewErrorToConsoleHook())

	logger.Debugf("Debug %d", 0)
	logger.Infof("Info %d", 1)
	logger.Warnf("Warn %d", 2)
	logger.Errorf("Error %d", 3)
	//logger.Panicf("Panic %d", 4)
}
