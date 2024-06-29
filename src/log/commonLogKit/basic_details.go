package commonLogKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v3/src/core/memoryKit"
	"github.com/richelieu-yang/chimera/v3/src/core/osKit"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/core/runtimeKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/userKit"
	"github.com/richelieu-yang/chimera/v3/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v3/src/diskKit"
	"github.com/richelieu-yang/chimera/v3/src/ip/ipKit"
	"github.com/richelieu-yang/chimera/v3/src/processKit"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"github.com/richelieu-yang/chimera/v3/src/time/timeKit"
	"strings"
)

func PrintBasicDetails(logger CommonLogger) {
	if logger == nil {
		return
	}

	logger.Info(strings.Repeat("=", 42))
	logger.Infof("\n%s", consts.Banner)

	logger.Infof("[CHIMERA, PROCESS] pid: [%d]", processKit.PID)

	/* golang */
	logger.Infof("[CHIMERA, GO] version: [%s]", runtimeKit.GoVersion)
	logger.Infof("[CHIMERA, GO] GOROOT: [%s]", runtimeKit.GoRoot)

	/* os */
	printOsDetails(logger)

	/* user */
	logger.Infof("[CHIMERA, USER] uid: [%s]", userKit.GetUid())
	logger.Infof("[CHIMERA, USER] gid: [%s]", userKit.GetGid())
	logger.Infof("[CHIMERA, USER] name: [%s]", userKit.GetName())
	logger.Infof("[CHIMERA, USER] user name: [%s]", userKit.GetUserName())
	logger.Infof("[CHIMERA, USER] home dir: [%s]", userKit.GetUserHomeDir())

	/* path */
	logger.Infof("[CHIMERA, PATH] working directory: [%s]", pathKit.GetWorkingDir())
	logger.Infof("[CHIMERA, PATH] os temporary directory: [%s]", pathKit.GetOsTempDir())
	logger.Infof("[CHIMERA, PATH] self dir: [%s]", pathKit.SelfDir())
	logger.Infof("[CHIMERA, PATH] main pkg path: [%s]", pathKit.MainPkgPath())

	/* json */
	logger.Infof("[CHIMERA, JSON] library: [%s]", jsonKit.GetLibrary())

	/* time */
	machineTime := timeKit.GetMachineTime()
	zoneName, zoneOffset := machineTime.Zone()
	logger.Infof("[CHIMERA, TIME] machine time: [%v], zone: [%s, %d]", machineTime, zoneName, zoneOffset)
	// Richelieu: 先注释掉，以防（断网或弱网环境下）导致拖延服务启动3s
	//if networkTime, source, err := timeKit.GetNetworkTime(); err != nil {
	//	logger.WithError(err).Warn("[CHIMERA, TIME] fail to get network time")
	//} else {
	//	logger.Infof("[CHIMERA, TIME] network time: [%v], source: [%s]", networkTime, source)
	//}

	/* ip */
	logger.Infof("[CHIMERA, IP] internal ip: [%s]", ipKit.GetInternalIp())
	ips := ipKit.GetIps()
	logger.Infof("[CHIMERA, IP] ips: [%s]", sliceKit.Join(ips, ", "))

	/* host */
	if hostInfo, err := runtimeKit.GetHostInfo(); err != nil {
		logger.Warnf("[CHIMERA, HOST] fail to get host info, error: %s", err.Error())
	} else {
		logger.Infof("[CHIMERA, HOST] host name: [%s]", hostInfo.Hostname)
	}

	/* CPU */
	printCpuDetails(logger)

	/* memory */
	printMemoryDetails(logger)

	/* disk */
	printDiskDetails(logger)

	logger.Info(strings.Repeat("=", 42))
}

func printOsDetails(logger CommonLogger) {
	logger.Infof("[CHIMERA, OS] os: [%s]", osKit.OS)
	logger.Infof("[CHIMERA, OS] arch: [%s]", osKit.ARCH)
	logger.Infof("[CHIMERA, OS] bits: [%d]", osKit.GetOsBits())

	if str, err := osKit.GetUlimitInfo(); err != nil {
		logger.Warnf("[CHIMERA, OS] fail to get ulimit information, error: %s", err.Error())
	} else {
		logger.Infof("[CHIMERA, OS] ulimit information:\n%s", str)
	}

	if i, err := osKit.GetThreadsMax(); err != nil {
		logger.Warnf("[CHIMERA, OS] fail to get kernel.threads-max, error: %s", err.Error())
	} else {
		logger.Infof("[CHIMERA, OS] kernel.threads-max: [%d]", i)
	}
	if i, err := osKit.GetPidMax(); err != nil {
		logger.Warnf("[CHIMERA, OS] fail to get kernel.pid_max, error: %s", err.Error())
	} else {
		logger.Infof("[CHIMERA, OS] kernel.pid_max: [%d]", i)
	}
	if i, err := osKit.GetMaxMapCount(); err != nil {
		logger.Warnf("[CHIMERA, OS] fail to get vm.max_map_count, error: %s", err.Error())
	} else {
		logger.Infof("[CHIMERA, OS] vm.max_map_count: [%d]", i)
	}
}

func printMemoryDetails(logger CommonLogger) {
	stats, err := memoryKit.GetMachineMemoryStats()
	if err != nil {
		logger.Errorf("[CHIMERA, MEMORY] fail to get machine memory stats, error: %s", err.Error())
		return
	}
	str := fmt.Sprintf("total: %s, available: %s, used: %s, free: %s, used percent: %.2f%%",
		dataSizeKit.ToReadableIecString(float64(stats.Total)),
		dataSizeKit.ToReadableIecString(float64(stats.Available)),
		dataSizeKit.ToReadableIecString(float64(stats.Used)),
		dataSizeKit.ToReadableIecString(float64(stats.Free)),
		stats.UsedPercent,
	)
	logger.Infof("[CHIMERA, MEMORY] machine memory stats: [%s]", str)
}

func printDiskDetails(logger CommonLogger) {
	stats, err := diskKit.GetDiskUsageStats()
	if err != nil {
		logger.Warnf("[CHIMERA, DISK] fail to get disk usage stats, error: %s", err.Error())
	} else {
		str := fmt.Sprintf("path: %s, free: %s, used: %s, total: %s, used percent: %.2f%%",
			stats.Path,
			dataSizeKit.ToReadableIecString(float64(stats.Free)),
			dataSizeKit.ToReadableIecString(float64(stats.Used)),
			dataSizeKit.ToReadableIecString(float64(stats.Total)),
			stats.UsedPercent,
		)
		logger.Infof("[CHIMERA, DISK] disk usage stats: [%s]", str)
	}
}

func printCpuDetails(logger CommonLogger) {
	logger.Infof("[CHIMERA, CPU] in a virtual machine? [%t]", cpuKit.InVirtualMachine())
	logger.Infof("[CHIMERA, CPU] vendor id: [%s]", cpuKit.GetVendorID())
	logger.Infof("[CHIMERA, CPU] vendor string: [%s]", cpuKit.GetVendorString())
	logger.Infof("[CHIMERA, CPU] brand name: [%s]", cpuKit.GetBrandName())
	logger.Infof("[CHIMERA, CPU] CPU number: [%d]", cpuKit.GetCpuNumber())
	logger.Infof("[CHIMERA, CPU] features: [%s]", sliceKit.Join(cpuKit.GetFeatureSet(), ","))
	logger.Infof("[CHIMERA, CPU] frequency: [%d]hz", cpuKit.GetFrequency())

	usage, err := cpuKit.GetUsagePercent()
	if err != nil {
		logger.Warnf("[CHIMERA, CPU] fail to get uasge percent, error: %s", err.Error())
	} else {
		logger.Infof("[CHIMERA, CPU] uasge percent: [%.2f]%%", usage)
	}
}
