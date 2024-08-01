package slbKit

import "github.com/richelieu-yang/chimera/v3/src/core/errorKit"

var (
	HaveNotBeenStartedError = errorKit.Newf("have not been started")

	AlreadyStartedError = errorKit.Newf("already started")

	AlreadyDisposedError = errorKit.Newf("already disposed")
)
