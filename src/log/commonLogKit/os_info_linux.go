package commonLogKit

import (
	"bufio"
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/cmd/cmdKit"
	"github.com/richelieu-yang/chimera/v3/src/core/osKit"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"os"
	"os/exec"
	"path/filepath"
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

	// getKernelParameterValue 获取Linux内核参数的值.
	var getKernelParameterValue = func(key string) (string, error) {
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

	for _, key := range kernelParameterKeys {
		value, err := getKernelParameterValue(key)
		if err != nil {
			logger.Warnf("[CHIMERA, OS] %s: fail to get, error: %s", key, err.Error())
			continue
		}
		logger.Infof("[CHIMERA, OS] %s: %s", key, value)
	}
}

// printCgroupInfo
/*
命令（stat -fc %T /sys/fs/cgroup/）用于判断系统当前使用的是 cgroup v1 还是 cgroup v2:
(1) 如果输出为 tmpfs，则表示系统使用的是 cgroup v1;
(2) 如果输出为 cgroup2fs，则表示系统使用的是 cgroup v2.
*/
func printCgroupInfo(logger Logger) {
	cgroupType, err := cmdKit.ExecuteToString(context.TODO(), "bash", "-c", "stat -fc %T /sys/fs/cgroup/")
	if err != nil {
		logger.Warnf("Fail to get cgroup type, error: %s", err.Error())
		return
	}
	logger.Infof("cgroup type: %s", cgroupType)

	var hardLimitPath, softLimitPath string
	if cgroupType == "cgroup2fs" {
		/* cgroup v2 */
		cgroupPath, err := getCgroupPath()
		if err != nil {
			logger.Warnf("Fail to get cgroup path, error: %s", err.Error())
			return
		}
		softLimitPath = pathKit.Join(cgroupPath, "memory.min") // v2 中没有直接对应的软限制
		hardLimitPath = pathKit.Join(cgroupPath, "memory.max")
	} else {
		/* cgroup v1 */
		softLimitPath = "/sys/fs/cgroup/memory/memory.soft_limit_in_bytes"
		hardLimitPath = "/sys/fs/cgroup/memory/memory.limit_in_bytes"
	}

	printLine := func(path string) {
		s := strKit.Split(path, osKit.PathSeparator)

		logger.Debugf("[TEST] len(s): %d", len(s))

		if len(s) <= 1 {
			return
		}
		key := s[len(s)-1]

		cmd := exec.Command("bash", "-c", fmt.Sprintf("cat %s", hardLimitPath))

		logger.Debugf("[TEST] command: %s", cmd.String())

		data, err := cmd.CombinedOutput()
		if err != nil {
			logger.Warnf("Command(%s) fails, error: %s", cmd.String(), err.Error())
			return
		}
		value := strKit.TrimSpace(string(data))
		logger.Infof("%s: %s", key, value)
	}
	printLine(softLimitPath)
	printLine(hardLimitPath)
}

// getCgroupPath 获取当前进程的 cgroup 路径（cgroup v2）
func getCgroupPath() (string, error) {
	file, err := os.Open("/proc/self/cgroup")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// cgroup2 的格式为 0::/path/to/cgroup
		parts := strings.Split(line, ":")
		if len(parts) == 3 && parts[0] == "0" {
			// 获取 cgroup 路径
			return filepath.Join("/sys/fs/cgroup", parts[2]), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("cgroup path not found")
}
