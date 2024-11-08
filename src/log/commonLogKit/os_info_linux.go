package commonLogKit

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"os/exec"
	"strings"
)

func printUlimitInformation(logger Logger) {
	/*
		Richelieu: 针对 Ubuntu，使用 bash 而 非sh，因为默认情况下 /bin/sh -> dash*.
	*/
	//cmd := exec.CommandContext(context.TODO(), "sh", "-c", "ulimit -a")
	cmd := exec.CommandContext(context.TODO(), "bash", "-c", "ulimit -a")
	data, err := cmd.CombinedOutput()
	if err != nil {
		logger.Warnf("[CHIMERA, OS] command(%s) fails, error: %s", cmd.String(), err.Error())
		return
	}
	str := strings.TrimSpace(string(data))
	logger.Infof("[CHIMERA, OS] ulimit imformation(ulimit -a):\n%s", str)
}

func printOsInformation(logger Logger) {
	var kernelParameterKeys = []string{
		"fs.aio-max-nr",
		"fs.file-max",
		"fs.inotify.max_user_watches",
		"fs.nr_open",
		"vm.dirty_background_ratio",
		"vm.dirty_ratio",
		"vm.swappiness",
		"vm.vfs_cache_pressure",
		"kernel.pid_max",
		"kernel.threads-max",
	}

	for _, key := range kernelParameterKeys {
		value, err := getKernelParameterValue(key)
		if err != nil {
			logger.Warnf("[CHIMERA, OS] Fail to get kernel parameter(%s), error: %s", key, err.Error())
			continue
		}
		logger.Infof("[CHIMERA, OS] %s: %s", key, value)
	}
}

// getKernelParameterValue 获取Linux内核参数的值.
func getKernelParameterValue(key string) (string, error) {
	s := strKit.Split(key, ".")
	path := pathKit.Join("/proc/sys", strKit.Join(s, "/"))

	cmd := exec.CommandContext(context.TODO(), "bash", "-c", fmt.Sprintf("cat %s", path))
	data, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	str := strings.TrimSpace(string(data))
	return str, nil
}
