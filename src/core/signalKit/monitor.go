package signalKit

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

// MonitorExitSignals 监听退出信号（拦截关闭信号）.
/*
可以参考 go-zero 中的 "core/proc/signals.go".

@param callback 可以为nil

PS:
(1) 无法拦截部分信号（e.g. syscall.SIGSTOP、syscall.SIGKILL）;
(2) 可以通过 logrus.RegisterExitHandler() 在程序退出前"毁尸灭迹"（在里面你甚至可以 time.Sleep）;
(3) 此函数对 主动调用os.Exit() 无效;
(4) 信号处理函数中不要使用 fmt.Println 等函数，因为它们不是线程安全的，会导致程序崩溃;
(5) 虽然可以多次调用本函数，但不推荐这么干，1次就够了.
*/
func MonitorExitSignals(callback func(sig os.Signal)) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, ExitSignals...)

	go func() {
		sig := <-ch

		if callback != nil {
			callback(sig)
		}

		logrus.WithField("signal", sig.String()).Fatal("Receive an exit signal.")
	}()
}
