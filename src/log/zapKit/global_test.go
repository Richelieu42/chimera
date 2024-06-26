package zapKit

import (
	"go.uber.org/zap"
	"testing"
)

func TestDebug(t *testing.T) {
	Debug("hello world", zap.String("key", "value"), zap.Bool("flag", true))

	Debugf("hello %s", "world")
	Debugw("hello world", "key", "value", "flag", true)
	Debugln("hello", "world")
}
