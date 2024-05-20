package main

import (
	//_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	textFormatter := logrusKit.NewDefaultTextFormatter()
	textFormatter.PadLevelText = false
	textFormatter.DisableLevelTruncation = false
	logrus.SetFormatter(textFormatter)

	logrus.WithField("1", "2").Debug("Debug")
	logrus.Info("Info")
	logrus.Warn("Warn")
	logrus.Error("Error")
}
