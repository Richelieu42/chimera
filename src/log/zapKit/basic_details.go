package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/log/commonLogKit"
	"go.uber.org/zap"
)

func PrintBasicDetails(loggers ...*zap.SugaredLogger) {
	var sl *zap.SugaredLogger
	loggers = sliceKit.RemoveZeroValues(loggers)
	if len(loggers) > 0 {
		sl = loggers[0]
	} else {
		sl = NewLogger(nil).Sugar()
		defer sl.Sync()
	}

	commonLogKit.PrintBasicDetails(sl)
}
