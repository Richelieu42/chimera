package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/richelieu-yang/chimera/v3/src/core/signalKit"
	"os"
)

func main() {
	signalKit.MonitorExitSignals(func(sig os.Signal) {
		time.Sleep(time.Second)
		logrus.Info(sig.String())
	})
	signalKit.MonitorExitSignals(func(sig os.Signal) {
		time.Sleep(time.Second)
		logrus.Info(sig.String())
	})
	signalKit.MonitorExitSignals(func(sig os.Signal) {
		time.Sleep(time.Second)
		logrus.Info(sig.String())
	})
	select {}
}
