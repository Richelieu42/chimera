package zapKit

import "testing"

func TestDebug(t *testing.T) {
	Debugf("hello %s", "world")
	Debugw("hello world", "key", "value", "flag", true)
	Debugln("hello", "world")
}
