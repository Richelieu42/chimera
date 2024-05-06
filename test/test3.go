package main

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
)

func main() {
	ctx := context.TODO()

	fmt.Println(ctx.Err() == nil)                               // true
	fmt.Println(errors.Is(ctx.Err(), context.Canceled))         // false
	fmt.Println(errors.Is(ctx.Err(), context.DeadlineExceeded)) // false
}

//func main() {
//	//signalKit.MonitorExitSignals(func(sig os.Signal) {
//	//	time.Sleep(time.Second)
//	//	logrus.Info(sig.String())
//	//})
//	//signalKit.MonitorExitSignals(func(sig os.Signal) {
//	//	time.Sleep(time.Second)
//	//	logrus.Info(sig.String())
//	//})
//	//signalKit.MonitorExitSignals(func(sig os.Signal) {
//	//	time.Sleep(time.Second)
//	//	logrus.Info(sig.String())
//	//})
//	//select {}
//}
