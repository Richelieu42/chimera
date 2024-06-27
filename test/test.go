package main

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
)

func main() {
	zapcore.RegisterHooks()

	zapcore.SamplerHook()

	logrus.RegisterExitHandler()

}
