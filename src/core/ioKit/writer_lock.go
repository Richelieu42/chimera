package ioKit

import (
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

// NewLockedWriter
/*
PS:
(1) zapcore.WriteSyncer 接口继承了 io.Writer 接口.
(2) 内部做了处理，可以多次调用（虽然这么干不推荐也没意义）.

@return 支持并发 Write
*/
func NewLockedWriter(w io.Writer) zapcore.WriteSyncer {
	if w == nil {
		return nil
	}

	switch w {
	case os.Stdout:
		return console.LockedWriteSyncerStdout
	case os.Stderr:
		return console.LockedWriteSyncerStderr
	default:
		return zapcore.Lock(zapcore.AddSync(w))
	}
}
