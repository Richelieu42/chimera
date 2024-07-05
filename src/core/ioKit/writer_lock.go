package ioKit

import (
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var (
	// LockedWriteSyncerStdout （加锁的）标准输出.
	LockedWriteSyncerStdout = zapcore.Lock(os.Stdout)

	// LockedWriteSyncerStderr （加锁的）标准错误输出.
	LockedWriteSyncerStderr = zapcore.Lock(os.Stderr)
)

// NewLockedWriteSyncer
/*
PS:
(1) zapcore.WriteSyncer 接口继承了 io.Writer 接口.
(2) 内部做了处理，可以多次调用（虽然这么干不推荐也没意义）.

@return 支持并发 Write
*/
func NewLockedWriteSyncer(w io.Writer) zapcore.WriteSyncer {
	if w == nil {
		return nil
	}

	switch w {
	case os.Stdout:
		return LockedWriteSyncerStdout
	case os.Stderr:
		return LockedWriteSyncerStderr
	default:
		return zapcore.Lock(zapcore.AddSync(w))
	}
}
