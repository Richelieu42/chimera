package appKit

import (
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"time"
)

var (
	Exit func(codes ...int) = zapKit.Exit

	RegisterExitHandler func(handlers ...func()) = zapKit.RegisterExitHandler

	RegisterParallelExitHandler func(handlers ...func()) = zapKit.RegisterParallelExitHandler

	// SetExitTimeout 执行所有exit handler的超时时间.
	SetExitTimeout func(d time.Duration) = zapKit.SetExitTimeout

	RunExitHandlers func() = zapKit.RunExitHandlers
)
