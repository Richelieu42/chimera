package main

import "go.uber.org/zap/zapcore"

func main() {
	zapcore.NewCore()
	zapcore.NewNopCore()
	zapcore.NewLazyWith()
	zapcore.NewSamplerWithOptions()
	zapcore.NewTee()
}
