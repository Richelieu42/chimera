package timeKit

import (
	"context"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestSetInterval(t *testing.T) {
	ctx := context.TODO()
	i := SetInterval(ctx, func(t time.Time) {
		logrus.Info(t)
	}, time.Second)

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 3)
	logrus.Info("sleep ends")

	ClearInterval(i)
	ClearInterval(i)

	logrus.Info("sleep1 starts")
	time.Sleep(time.Second * 3)
	logrus.Info("sleep1 ends")
}

func TestSetInterval1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Millisecond*1500)
	defer cancel()

	i := SetInterval(ctx, func(t time.Time) {
		logrus.Info(t)
	}, time.Second)

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 3)
	logrus.Info("sleep ends")

	ClearInterval(i)
	ClearInterval(i)

	logrus.Info("sleep1 starts")
	time.Sleep(time.Second * 3)
	logrus.Info("sleep1 ends")
}
