package pushKit

import "github.com/richelieu-yang/chimera/v3/src/core/errorKit"

var (
	NotSetupError = errorKit.Newf("haven’t been set up correctly")

	ChannelClosedError = errorKit.Newf("channel has already been closed")

	NoSuitableChannelError = errorKit.Newf("no suitable channel")
)
