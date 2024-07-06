package statKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v3/src/core/mathKit"
	"github.com/richelieu-yang/chimera/v3/src/core/memoryKit"
	"github.com/richelieu-yang/chimera/v3/src/core/osKit"
	"github.com/richelieu-yang/chimera/v3/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v3/src/diskKit"
	"github.com/richelieu-yang/chimera/v3/src/processKit"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"go.uber.org/zap"
	"os"
	"runtime"
	"sync"
)

type (
	Stats struct {
		Program *ProgramStats `json:"program"`

		Machine *MachineStats `json:"machine"`
	}

	ProgramStats struct {
		PID            int `json:"pid"`
		GoroutineCount int `json:"goroutineCount"`

		CpuUsagePercent      float64 `json:"cpuUsagePercent"`
		CpuUsagePercentError error   `json:"cpuUsagePercentError,omitempty"`

		Memory *ProgramMemoryStats `json:"memory"`
	}

	MachineStats struct {
		CpuUsagePercent      float64 `json:"cpuUsagePercent"`
		CpuUsagePercentError error   `json:"cpuUsagePercentError,omitempty"`

		DiskPath              string  `json:"diskPath,omitempty"`
		DiskUsagePercent      float64 `json:"diskUsagePercent,omitempty"`
		DiskUsagePercentError error   `json:"diskUsagePercentError,omitempty"`

		// ProcessCount 进程数
		ProcessCount      int   `json:"processCount,omitempty"`
		ProcessCountError error `json:"processCountError,omitempty"`

		// ProcessThreadCount 进程数（包括线程数）
		ProcessThreadCount      int   `json:"processThreadCount,omitempty"`
		ProcessThreadCountError error `json:"processThreadCountError,omitempty"`

		MaxProcessThreadCountByUser      int    `json:"maxProcessThreadCountByUser,omitempty"`
		MaxProcessThreadCountByUserError string `json:"maxProcessThreadCountByUserError,omitempty"`
		PidMax                           int    `json:"pidMax,omitempty"`
		PidMaxError                      string `json:"pidMaxError,omitempty"`
		ThreadsMax                       int    `json:"threadsMax,omitempty"`
		ThreadsMaxError                  string `json:"threadsMaxError,omitempty"`
		MaxMapCount                      int    `json:"maxMapCount,omitempty"`
		MaxMapCountError                 string `json:"maxMapCountError,omitempty"`

		Memory *MachineMemoryStats `json:"memory"`
	}

	ProgramMemoryStats struct {
		Alloc      string `json:"alloc"`
		TotalAlloc string `json:"totalAlloc"`
		Sys        string `json:"sys"`
		NumGC      uint32 `json:"numGC"`
		EnableGC   bool   `json:"enableGC"`
	}

	MachineMemoryStats struct {
		Error       string  `json:"error,omitempty"`
		Total       string  `json:"total,omitempty"`
		Available   string  `json:"available,omitempty"`
		Used        string  `json:"used,omitempty"`
		UsedPercent float64 `json:"usedPercent,omitempty"`
		Free        string  `json:"free,omitempty"`
	}
)

// GetStats
/*
PS: 由于获取CPU使用率耗时较长，本函数内部使用 sync.WaitGroup.
*/
func GetStats() *Stats {
	rst := &Stats{
		Program: &ProgramStats{
			Memory: &ProgramMemoryStats{},
		},
		Machine: &MachineStats{
			Memory: &MachineMemoryStats{},
		},
	}
	programStats := rst.Program
	machineStats := rst.Machine

	var wg sync.WaitGroup

	/* program */
	wg.Add(1)
	go func() {
		defer wg.Done()

		programStats.PID = os.Getpid()
		programStats.GoroutineCount = runtime.NumGoroutine()

		{
			stats := memoryKit.GetProgramMemoryStats()
			programStats.Memory.Alloc = dataSizeKit.ToReadableIecString(float64(stats.Alloc))
			programStats.Memory.TotalAlloc = dataSizeKit.ToReadableIecString(float64(stats.TotalAlloc))
			programStats.Memory.Sys = dataSizeKit.ToReadableIecString(float64(stats.Sys))
			programStats.Memory.NumGC = stats.NumGC
			programStats.Memory.EnableGC = stats.EnableGC
		}

		if usagePercent, err := cpuKit.GetProcessUsagePercent(int32(programStats.PID)); err != nil {
			programStats.CpuUsagePercentError = err
		} else {
			programStats.CpuUsagePercent = mathKit.Round(usagePercent, 2)
		}
	}()

	/* machine */
	// (1) CPU
	wg.Add(1)
	go func() {
		defer wg.Done()

		usagePercent, err := cpuKit.GetUsagePercent()
		if err != nil {
			machineStats.CpuUsagePercentError = err
		} else {
			machineStats.CpuUsagePercent = mathKit.Round(usagePercent, 2)
		}
	}()

	// (2) disk
	wg.Add(1)
	go func() {
		defer wg.Done()

		stats, err := diskKit.GetDiskUsageStats()
		if err != nil {
			machineStats.DiskUsagePercentError = err
		} else {
			machineStats.DiskPath = stats.Path
			machineStats.DiskUsagePercent = mathKit.Round(stats.UsedPercent, 2)
		}
	}()

	// (3) others
	wg.Add(1)
	go func() {
		defer wg.Done()

		processCount, err := processKit.GetProcessCount()
		if err != nil {
			machineStats.ProcessCountError = err
		} else {
			machineStats.ProcessCount = processCount
		}

		processThreadCount, err := processKit.GetProcessThreadCount()
		if err != nil {
			machineStats.ProcessThreadCountError = err
		} else {
			machineStats.ProcessThreadCount = processThreadCount
		}

		// memory
		{
			stats, err := memoryKit.GetMachineMemoryStats()
			if err != nil {
				machineStats.Memory.Error = err.Error()
			} else {
				machineStats.Memory.Total = dataSizeKit.ToReadableIecString(float64(stats.Total))
				machineStats.Memory.Available = dataSizeKit.ToReadableIecString(float64(stats.Available))
				machineStats.Memory.Used = dataSizeKit.ToReadableIecString(float64(stats.Used))
				machineStats.Memory.UsedPercent = mathKit.Round(stats.UsedPercent, 2)
				machineStats.Memory.Free = dataSizeKit.ToReadableIecString(float64(stats.Free))
			}
		}

		// ulimit -u
		if tmp, err := osKit.GetMaxProcessThreadCountByUser(); err != nil {
			machineStats.MaxProcessThreadCountByUserError = err.Error()
		} else {
			machineStats.MaxProcessThreadCountByUser = tmp
		}
		// kernel.pid_max
		if tmp, err := osKit.GetPidMax(); err != nil {
			machineStats.PidMaxError = err.Error()
		} else {
			machineStats.PidMax = tmp
		}
		// kernel.threads-max
		if tmp, err := osKit.GetThreadsMax(); err != nil {
			machineStats.ThreadsMaxError = err.Error()
		} else {
			machineStats.ThreadsMax = tmp
		}
		// vm.max_map_count
		if tmp, err := osKit.GetMaxMapCount(); err != nil {
			machineStats.MaxMapCountError = err.Error()
		} else {
			machineStats.MaxMapCount = tmp
		}
	}()

	wg.Wait()
	return rst
}

func PrintStats(logger *zap.SugaredLogger) {
	stats := GetStats()
	json, err := jsonKit.MarshalIndentToString(stats, "", "    ")
	if err != nil {
		logger.Errorf("Fail to marshal, error: %s", err.Error())
		return
	}
	logger.Infof("[CHIMERA] stats:\n%s", json)
}
