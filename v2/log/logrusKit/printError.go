package logrusKit

import (
	"github.com/richelieu42/chimera/v2/core/errorKit"
	"github.com/sirupsen/logrus"
)

// PrintError
/*
@param err 只有通过"github.com/pkg/errors" New 或 Wrap 的error才有堆栈信息
*/
func PrintError(err error) {
	PrintErrorWithLogger(err, nil)
}

func PrintErrorWithLogger(err error, logger *logrus.Logger) {
	if err == nil {
		return
	}
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	// 输出
	if errorKit.IsErrorWithStack(err) {
		cause := errorKit.Cause(err)
		logger.Errorf("%T %v", cause, cause)
		logger.Errorf("%+v", err)
		return
	}
	logger.Errorf("%v", err)
}
