package console

import (
	"go.uber.org/zap"
	"testing"
)

func TestDebug(t *testing.T) {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
}

func TestDebug1(t *testing.T) {
	defer Sync()

	Debug("ddd")

	/* WrappedLogger */
	Debug("hello world", zap.String("key", "value"), zap.Bool("flag", true))

	/* Sugar WrappedLogger */
	Debugf("hello %s", "world")
	Debugw("hello world", "key", "value", "flag", true)
	Debugln("hello", "world")

	Info("Info")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal")
}

func TestSetDefaultLevel(t *testing.T) {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")

	SetDefaultLevel(zap.InfoLevel)
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
}
