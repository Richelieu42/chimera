package commonLogKit

import (
	"context"
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
	"time"
)

func PrintBasicDetails(logger Logger) {
	if logger == nil {
		return
	}

	logger.Info(strings.Repeat("=", 42))
	logger.Infof("\n%s", consts.Banner)

	chA := make(chan struct{})
	chB := make(chan struct{})

	/* 协程a */
	go func() {
		logger.Infof("[CHIMERA, PROCESS] pid: [%d]", processKit.PID)

		/* golang */
		logger.Infof("[CHIMERA, GO] version: [%s]", runtimeKit.GetGoVersion())
		logger.Infof("[CHIMERA, GO] GOROOT: [%s]", runtimeKit.GetGoRoot())

		/* os */
		logger.Infof("[CHIMERA, OS] os: [%s]", osKit.OS)
		logger.Infof("[CHIMERA, OS] arch: [%s]", osKit.ARCH)
		logger.Infof("[CHIMERA, OS] bits: [%d]", osKit.GetOsBits())
		printUlimitInformation(logger)
		printOsInformation(logger)

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

		// 关闭信道，通知协程a可以继续执行
		close(chA)
	}()

	/* 协程b */
	go func() {
		/* time */
		printTimeDetails(logger, chA)
	}()

	// a执行完毕
	<-chA
	select {
	case <-chB:
		// b执行完毕
	case <-time.After(time.Millisecond * 100):
		// 等了100ms（但b还在执行）
	}

	logger.Info(strings.Repeat("=", 42))
}

func printTimeDetails(logger Logger, ch chan struct{}) {
	reqCtx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	networkTime, source, err := timeKit.GetNetworkTime(reqCtx)
	machineTime := timeKit.GetMachineTime()
	zoneName, zoneOffset := machineTime.Zone()

	// 等待协程a执行完毕，为防止: 多协程输出导致输出混在一起
	<-ch

	if err != nil {
		logger.Warnf("[CHIMERA, TIME] Fail to get network time, error: %s", err.Error())
	} else {
		logger.Infof("[CHIMERA, TIME] network time: [%v], source: [%s]", networkTime, source)
	}
	logger.Infof("[CHIMERA, TIME] machine time: [%v], zone: [%s, %d]", machineTime, zoneName, zoneOffset)
}

func printMemoryDetails(logger Logger) {
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

func printDiskDetails(logger Logger) {
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

func printCpuDetails(logger Logger) {
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
