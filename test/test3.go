package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 0
		ch <- 1
		ch <- 2
		time.Sleep(time.Second * 3)
		close(ch)
	}()

	logrus.Info("start ---")
outerLoop:
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				break outerLoop
			}
			logrus.Info(i)
		}
	}
	logrus.Info("end ---")
}
