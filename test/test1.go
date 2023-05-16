package main

import (
	"github.com/richelieu42/chimera/v2/src/confKit"
	"github.com/richelieu42/chimera/v2/src/database/redisKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	path := "/Users/richelieu/GolandProjects/chimera/chimera-lib/config.yaml"

	logrusKit.MustSetUp(nil)

	type config struct {
		Redis *redisKit.Config `json:"redis"`
	}
	c := &config{}
	confKit.MustLoad(path, c)
	redisKit.MustSetUp(c.Redis)

	client, err := redisKit.GetClient()
	if err != nil {
		logrus.Fatal(err)
	}

	//go func() {
	//	time.Sleep(time.Millisecond * 100)
	//
	//	mu := client.NewDistributedMutex("/ccc" /*, redsync.WithRetryDelay(time.Second)*/)
	//	logrus.Info("[GOROUTINE] ready to lock on")
	//	if err := mu.Lock(); err != nil {
	//		logrus.WithFields(logrus.Fields{
	//			"error": err.Error(),
	//		}).Fatal("[GOROUTINE] fail to lock")
	//	}
	//	logrus.Info("[GOROUTINE] lock on")
	//
	//	ok, err := mu.Unlock()
	//	logrus.WithFields(logrus.Fields{
	//		"ok":  ok,
	//		"err": err,
	//	}).Info("[GOROUTINE] lock off")
	//}()

	mu := client.NewDistributedMutex("/ccc")
	logrus.Info("ready to lock on")
	if err := mu.Lock(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("lock on")

	select {}
}
