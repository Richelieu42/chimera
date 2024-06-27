package zapKit

import (
	"go.uber.org/zap"
	"io"
)

// WrapLogger
/*
适用场景:
*/
func WrapLogger(logger *zap.Logger, writers ...io.Writer) *WrappedLogger {
	return &WrappedLogger{
		Logger:  logger,
		Writers: writers,
	}
}

// WrapSugarLogger
/*
适用场景:
*/
func WrapSugarLogger(sugaredLogger *zap.SugaredLogger, writers ...io.Writer) *WrappedSugaredLogger {
	return &WrappedSugaredLogger{
		SugaredLogger: sugaredLogger,
		Writers:       writers,
	}
}
