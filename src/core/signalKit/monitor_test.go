package signalKit

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
)

func TestMonitorExitSignalsSynchronously(t *testing.T) {
	logrus.RegisterExitHandler(func() {
		logrus.Info("sleep1 starts...")
		time.Sleep(time.Second)
		logrus.Info("sleep1 ends...")
	})
	logrus.RegisterExitHandler(func() {
		logrus.Info("sleep2 starts...")
		time.Sleep(time.Second * 3)
		logrus.Info("sleep2 ends...")
	})

	MonitorExitSignalsSynchronously(func(sig os.Signal) {
		logrus.Info("sleep0 starts...")
		time.Sleep(time.Second)
		logrus.Info("sleep0 ends...")
	})
}
