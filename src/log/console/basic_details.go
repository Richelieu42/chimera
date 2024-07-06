package console

import (
	"github.com/richelieu-yang/chimera/v3/src/log/commonLogKit"
)

func PrintBasicDetails() {
	logger := S()
	defer logger.Sync()

	commonLogKit.PrintBasicDetails(logger)
}
