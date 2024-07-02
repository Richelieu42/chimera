package zapInitKit

import (
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
)

func init() {
	/* 替换 zap库 的全局loggers */
	l := zapKit.NewLogger(nil)
	zap.ReplaceGlobals(l)
}

// SetUp Deprecated: 使用"import _".
func SetUp() {
}
