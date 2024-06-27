package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	"go.uber.org/zap"
	"io"
)

type (
	WrappedLogger struct {
		*zap.Logger

		Writers []io.Writer
	}

	WrappedSugaredLogger struct {
		*zap.SugaredLogger

		Writers []io.Writer
	}
)

func (l *WrappedLogger) Close() (err error) {
	if l == nil {
		return
	}

	_ = l.Sync()
	for _, w := range l.Writers {
		tmpErr := ioKit.TryToClose(w)
		if tmpErr != nil && err == nil {
			err = tmpErr
		}
	}
	return
}

func (sl *WrappedSugaredLogger) Close() (err error) {
	if sl == nil {
		return
	}

	_ = sl.Sync()
	for _, w := range sl.Writers {
		tmpErr := ioKit.TryToClose(w)
		if tmpErr != nil && err == nil {
			err = tmpErr
		}
	}
	return
}
