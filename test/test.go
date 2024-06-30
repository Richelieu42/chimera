package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/osKit"
	"github.com/richelieu-yang/chimera/v3/src/core/signalKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 自定义的 Fatal Hook
type myFatalHook struct{}

func (h *myFatalHook) OnWrite(*zapcore.CheckedEntry, []zap.Field) {
	// 在这里实现你的钩子逻辑，比如执行清理工作
	fmt.Println("Fatal hook triggered!")
}

func main() {
	osKit.Exit()

	signalKit.MonitorExitSignals()
}

func a() {
	defer func() {
		fmt.Println(666)
	}()

	panic("000")
}
