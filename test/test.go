package main

import (
	"bytes"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap/zapcore"
)

func main() {
	out := &bytes.Buffer{}
	ws := zapcore.Lock(zapcore.AddSync(out))
	encoder := zapKit.NewEncoder()
	core := zapKit.NewCore(encoder, ws, zapcore.DebugLevel)
	logger := zapKit.NewLogger(core)

	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")

	_ = ws.Sync()

	console.Infof("content:\n%s\n", out.String())
	//fmt.Println(out.String())
}
