package statKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func MustSetup(logPath string) {
	if err := Setup(logPath); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func Setup(logPath string) error {
	if strKit.IsBlank(logPath) {
		// (1) 输出到: 控制台
	} else {
		// (2) 输出到: 文件日志
		if err := fileKit.AssertNotExistOrIsFile(logPath); err != nil {
			return err
		}
		f, err := fileKit.CreateInAppendMode(logPath)
		if err != nil {
			return err
		}
		logger = logrusKit.NewLogger(logrusKit.WithOutput(f))
	}

	c, _, err := cronKit.NewCronWithTask("@every 20s", func() {
		PrintStats(logger)
	})
	if err != nil {
		return err
	}
	c.Start()
	return nil
}
