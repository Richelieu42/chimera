package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/appKit"
	"go.uber.org/zap/zapcore"
)

// DefaultFatalHook 默认的 fatal hook
type DefaultFatalHook struct {
}

func (h *DefaultFatalHook) OnWrite(entry *zapcore.CheckedEntry, fields []zapcore.Field) {
	appKit.Exit(1)
}
