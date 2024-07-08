package zapKit

import (
	"testing"
)

func TestNewStdLogger(t *testing.T) {
	l := NewLogger(nil)
	l.Debug("Debug")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")

	logger := NewStdLogger(l)
	logger.Print("hello world")
	logger.Println("hello world")
}
