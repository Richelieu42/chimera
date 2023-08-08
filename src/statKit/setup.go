package statKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func MustSetup(logPath string) {
	if err := Setup(logPath); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func Setup(logPath string) error {
	if err := fileKit.AssertNotExistOrIsFile(logPath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(logPath); err != nil {
		return err
	}

	return nil
}
