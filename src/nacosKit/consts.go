package nacosKit

import "github.com/richelieu-yang/chimera/v2/src/core/ptrKit"

type (
	NacosLogLevel *string
)

var (
	DebugLevel = ptrKit.ToPtr("debug")
	InfoLevel  = ptrKit.ToPtr("info")
	WarnLevel  = ptrKit.ToPtr("warn")
	ErrorLevel = ptrKit.ToPtr("error")
)
