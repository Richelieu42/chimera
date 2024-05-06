package signalKit

import (
	"github.com/richelieu-yang/chimera/v3/src/processKit"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
)

func TestMonitorExitSignalsSynchronously(t *testing.T) {
	logrus.Infof("pid: [%d]", processKit.PID)

	logrus.RegisterExitHandler(func() {
		logrus.Info("sleep1 starts...")
		time.Sleep(time.Second)
		logrus.Info("sleep1 ends...")
	})
	logrus.RegisterExitHandler(func() {
		logrus.Info("sleep2 starts...")
		time.Sleep(time.Second)
		logrus.Info("sleep2 ends...")
	})

	MonitorExitSignalsSynchronously(func(sig os.Signal) {
		logrus.Info("sleep0 starts...")
		time.Sleep(time.Second)
		logrus.Info("sleep0 ends...")
	})
}
