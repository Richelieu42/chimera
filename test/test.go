package main

import "go.uber.org/zap/zapcore"

func main() {
	zapcore.NewNopCore()
}
