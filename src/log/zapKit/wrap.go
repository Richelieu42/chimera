package zapKit

import (
	"go.uber.org/zap"
	"io"
)

// WrapLogger
/*
适用场景: 使用完后，需要释放对应资源（关闭输出）.
*/
func WrapLogger(logger *zap.Logger, writers ...io.Writer) *WrappedLogger {
	return &WrappedLogger{
		Logger:  logger,
		Writers: writers,
	}
}

// WrapSugarLogger
/*
适用场景: 使用完后，需要释放对应资源（关闭输出）.
*/
func WrapSugarLogger(sugaredLogger *zap.SugaredLogger, writers ...io.Writer) *WrappedSugaredLogger {
	return &WrappedSugaredLogger{
		SugaredLogger: sugaredLogger,
		Writers:       writers,
	}
}
