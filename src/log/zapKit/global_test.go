package zapKit

import (
	"go.uber.org/zap"
	"testing"
)

func TestDebug(t *testing.T) {
	defer Sync()

	/* Logger */
	Debug("hello world", zap.String("key", "value"), zap.Bool("flag", true))

	/* Sugar Logger */
	Debugf("hello %s", "world")
	Debugw("hello world", "key", "value", "flag", true)
	Debugln("hello", "world")
}
