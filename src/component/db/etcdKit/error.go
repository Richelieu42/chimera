package etcdKit

import "github.com/richelieu-yang/chimera/v2/src/core/errorKit"

var (
	NotSetupError = errorKit.New("uninitialized component")
)
