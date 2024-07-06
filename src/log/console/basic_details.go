package console

import (
	"github.com/richelieu-yang/chimera/v3/src/log/commonLogKit"
)

func PrintBasicDetails() {
	defer s.Sync()

	commonLogKit.PrintBasicDetails(s)
}
