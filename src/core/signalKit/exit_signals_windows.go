package signalKit

import (
	"os"
	"syscall"
)

var (
	// ExitSignals 会将程序退出的所有能捕捉的信号
	/*
		参考: go-zero中的"core/proc/signals.go".
		2024/05/24: 应马瑞要求，添加: syscall.SIGINT
	*/
	ExitSignals = []os.Signal{syscall.SIGTERM, syscall.SIGINT}

	//// ExitSignals 会将程序退出的所有能捕捉的信号
	///*
	//	参考: https://mp.weixin.qq.com/s/ATiAnX0PrqkBVnSFBTiMGQ
	//*/
	//ExitSignals = []os.Signal{
	//	syscall.SIGHUP,
	//	//syscall.SIGUSR1,
	//	//syscall.SIGUSR2,
	//	syscall.SIGINT,
	//	syscall.SIGTERM,
	//	//syscall.SIGTSTP,
	//	syscall.SIGQUIT,
	//	//syscall.SIGSTOP,
	//	syscall.SIGKILL,
	//}
)
