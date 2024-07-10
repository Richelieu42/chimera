package zapKit

import (
	"log"
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

func TestRedirectStdLog(t *testing.T) {
	log.Println("123")

	l := NewLogger(nil)
	f := RedirectStdLog(l)
	f()
	log.Println("456")
}
