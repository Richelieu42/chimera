package zapKit

import "github.com/richelieu-yang/chimera/v3/src/core/osKit"

var (
	RegisterExitHandler func(handler func()) = osKit.RegisterExitHandler

	RegisterParallelExitHandler func(handler func()) = osKit.RegisterParallelExitHandler
)
