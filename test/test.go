package main

import (
	"bytes"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"go.uber.org/zap/zapcore"
)

func main() {
	out := &bytes.Buffer{}

	zapcore.WriteSyncer

	encoder := zapKit.NewEncoder()
	ws := zapcore.AddSync()
	core := zapKit.NewCore(encoder, ws, zapcore.DebugLevel)

	zapKit.NewCore()

	zapKit.NewLogger(nil)

	fmt.Print(netKit.JoinToHost("::1", 8888))
}

type ()
