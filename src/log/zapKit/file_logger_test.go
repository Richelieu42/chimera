package zapKit

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestNewFileLogger(t *testing.T) {
	path := "_file_logger.txt"

	l, err := NewFileLogger(path, "", zapcore.DebugLevel)
	if err != nil {
		panic(err)
	}
	l.Debug("Debug")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")
}
