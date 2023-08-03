package osKit

import "github.com/richelieu-yang/chimera/v2/src/core/errorKit"

func GetProcessCount() (int, error) {
	return 0, errorKit.New("not yet realized")
}

// GetThreadsMax 获取Linux的"kernel.threads-max"
func GetThreadsMax() (int, error) {
	return 0, errorKit.New("not yet realized")
}

// GetPidMax 获取Linux的"kernel.pid_max"
func GetPidMax() (int, error) {
	return 0, errorKit.New("not yet realized")
}
