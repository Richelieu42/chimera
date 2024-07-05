package console

import (
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	// LockedWriteSyncerStdout 标准输出.
	LockedWriteSyncerStdout = zapcore.Lock(os.Stdout)

	// LockedWriteSyncerStderr 标准错误输出.
	LockedWriteSyncerStderr = zapcore.Lock(os.Stderr)
)
