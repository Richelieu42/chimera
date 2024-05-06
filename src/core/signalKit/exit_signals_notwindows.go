//go:build !windows

package signalKit

import (
	"os"
	"syscall"
)

var (
	// ExitSignals 会将程序退出的所有能捕捉的信号
	/*
		参考: go-zero 中的"core/proc/signals.go".
	*/
	ExitSignals = []os.Signal{ /*syscall.SIGUSR1, syscall.SIGUSR2,*/ syscall.SIGTERM, syscall.SIGINT}

	//// ExitSignals 会将程序退出的所有能捕捉的信号
	///*
	//	参考: https://mp.weixin.qq.com/s/ATiAnX0PrqkBVnSFBTiMGQ
	//*/
	//ExitSignals = []os.Signal{
	//	syscall.SIGHUP,
	//	syscall.SIGUSR1,
	//	syscall.SIGUSR2,
	//	syscall.SIGINT,
	//	// e.g.kill ${pid} || kill -15 ${pid}
	//	syscall.SIGTERM,
	//	syscall.SIGTSTP,
	//	syscall.SIGQUIT,
	//
	//	// 无条件结束程序（不能被捕获、阻塞或忽略）
	//	syscall.SIGSTOP,
	//	// 停止进程（不能被捕获、阻塞或忽略）	e.g.kill -9 ${pid}
	//	syscall.SIGKILL,
	//}
)
