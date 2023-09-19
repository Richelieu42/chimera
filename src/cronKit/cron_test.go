package cronKit

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewCronWithTask(t *testing.T) {
	c, _, err := NewCronWithTask("@every 3s", func() {
		logrus.Info("do")
	})
	if err != nil {
		panic(err)
	}
	c.Start()

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 10)
	logrus.Info("sleep ends")

	// 可以多次调用 Cron.Stop()，虽然只有第一次有意义，但至少不会panic
	logrus.Info("0")
	select {
	case <-c.Stop().Done():
	}
	logrus.Info("1")
	select {
	case <-c.Stop().Done():
	}
	logrus.Info("2")

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 4)
	logrus.Info("sleep ends")
}
