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
	cancel := RedirectStdLog(l)

	log.Println("456")

	cancel()
	log.Println("789")
	/*
		Output:
		2024/07/10 14:03:55 123
		2024-07-10T14:03:55.580+0800	INFO	zapKit/std_logger_test.go:26	456
	*/
}
