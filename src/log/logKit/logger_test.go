package logKit

import (
	"testing"
)

func TestNewFileLogger(t *testing.T) {
	//logger, err := NewFileLogger("_test.log", "[TEST] ", 0644)
	logger, err := NewFileLogger("_test.log", "", 0644)
	if err != nil {
		panic(err)
	}

	logger.Println("0")
	logger.Println("1")
	logger.Println("2")
}
