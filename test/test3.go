package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		close(ch)
	}()

	logrus.Info("start ---")
outerLoop:
	for {
		select {
		case <-ch:
			logrus.Info("break")
			break outerLoop
		}
	}
	logrus.Info("end ---")
}
