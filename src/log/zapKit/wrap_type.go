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

func (wl *WrappedLogger) Close() (err error) {
	if wl == nil {
		return
	}

	_ = wl.Sync()
	for _, w := range wl.Writers {
		tmpErr := ioKit.TryToClose(w)
		if tmpErr != nil && err == nil {
			err = tmpErr
		}
	}
	return
}

func (wsl *WrappedSugaredLogger) Close() (err error) {
	if wsl == nil {
		return
	}

	_ = wsl.Sync()
	for _, w := range wsl.Writers {
		tmpErr := ioKit.TryToClose(w)
		if tmpErr != nil && err == nil {
			err = tmpErr
		}
	}
	return
}
