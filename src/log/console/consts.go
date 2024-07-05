package console

import (
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	// LockedWriteSyncerStdout （加锁的）标准输出.
	LockedWriteSyncerStdout = zapcore.Lock(os.Stdout)

	// LockedWriteSyncerStderr （加锁的）标准错误输出.
	LockedWriteSyncerStderr = zapcore.Lock(os.Stderr)
)
