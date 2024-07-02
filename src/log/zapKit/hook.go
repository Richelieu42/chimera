package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/appKit"
	"go.uber.org/zap/zapcore"
)

// defaultFatalHook 默认的 fatal hook
type defaultFatalHook struct {
}

func (h *defaultFatalHook) OnWrite(entry *zapcore.CheckedEntry, fields []zapcore.Field) {
	appKit.Exit(1)
}
