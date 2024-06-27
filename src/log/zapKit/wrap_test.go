package zapKit

import (
	"os"
	"testing"
)

func TestWrapLogger(t *testing.T) {
	f, err := os.Create("_a.log")
	if err != nil {
		panic(err)
	}

	logger := NewLogger(WithWriteSyncer())
	logger1 := WrapLogger()
}
