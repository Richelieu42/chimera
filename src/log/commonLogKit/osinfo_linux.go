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
		"fs.epoll.max_user_watches",
		"fs.file-max",
		"fs.inotify.max_user_watches",
		"fs.nr_open",

		"kernel.hung_task_timeout_secs",
		"kernel.pid_max",
		"kernel.sched_latency_ns",
		"kernel.sched_migration_cost_ns",
		"kernel.sched_min_granularity_ns",
		"kernel.threads-max",
		"kernel.timer_migration",
		"kernel.numa_balancing",

		//"memory.limit_in_bytes",
		//"memory.soft_limit_in_bytes",

		"net.core.rmem_max",
		"net.core.wmem_max",
		"net.core.somaxconn",
		"net.ipv4.ip_local_port_range",
		"net.ipv4.tcp_fin_timeout",
		"net.ipv4.tcp_max_syn_backlog",
		//"net.ipv4.tcp_tw_recycle", // 被 net.ipv4.tcp_tw_reuse 替代
		"net.ipv4.tcp_tw_reuse",

		"vm.dirty_background_ratio",
		"vm.dirty_ratio",
		"vm.max_map_count",
		"vm.overcommit_memory",
		"vm.swappiness",
		"vm.vfs_cache_pressure",
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
