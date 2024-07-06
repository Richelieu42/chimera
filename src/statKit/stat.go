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

		Alloc      string `json:"alloc"`
		TotalAlloc string `json:"totalAlloc"`
		Sys        string `json:"sys"`
		NumGC      uint32 `json:"numGC"`
		EnableGC   bool   `json:"enableGC"`

		CpuUsagePercent      float64 `json:"cpuUsagePercent"`
		CpuUsagePercentError error   `json:"cpuUsagePercentError,omitempty"`
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

		MemoryStatsError error   `json:"memoryStatsError,omitempty"`
		Total            string  `json:"total,omitempty"`
		Available        string  `json:"available,omitempty"`
		Used             string  `json:"used,omitempty"`
		UsedPercent      float64 `json:"usedPercent,omitempty"`
		Free             string  `json:"free,omitempty"`

		MaxProcessThreadCountByUser      int    `json:"maxProcessThreadCountByUser,omitempty"`
		MaxProcessThreadCountByUserError string `json:"maxProcessThreadCountByUserError,omitempty"`
		PidMax                           int    `json:"pidMax,omitempty"`
		PidMaxError                      string `json:"pidMaxError,omitempty"`
		ThreadsMax                       int    `json:"threadsMax,omitempty"`
		ThreadsMaxError                  string `json:"threadsMaxError,omitempty"`
		MaxMapCount                      int    `json:"maxMapCount,omitempty"`
		MaxMapCountError                 string `json:"maxMapCountError,omitempty"`
	}
)

// GetStats
/*
PS: 由于获取CPU使用率耗时较长，本函数内部使用 sync.WaitGroup.
*/
func GetStats() *Stats {
	pStats := &ProgramStats{}
	mStats := &MachineStats{}
	rst := &Stats{
		Program: pStats,
		Machine: mStats,
	}
	var wg sync.WaitGroup

	/* program */
	wg.Add(1)
	go func() {
		defer wg.Done()

		pStats.PID = os.Getpid()
		pStats.GoroutineCount = runtime.NumGoroutine()

		stats := memoryKit.GetProgramMemoryStats()
		pStats.Alloc = dataSizeKit.ToReadableIecString(float64(stats.Alloc))
		pStats.TotalAlloc = dataSizeKit.ToReadableIecString(float64(stats.TotalAlloc))
		pStats.Sys = dataSizeKit.ToReadableIecString(float64(stats.Sys))
		pStats.NumGC = stats.NumGC
		pStats.EnableGC = stats.EnableGC

		if usagePercent, err := cpuKit.GetProcessUsagePercent(int32(pStats.PID)); err != nil {
			pStats.CpuUsagePercentError = err
		} else {
			pStats.CpuUsagePercent = mathKit.Round(usagePercent, 2)
		}
	}()

	/* machine */
	// (1) CPU
	wg.Add(1)
	go func() {
		defer wg.Done()

		usagePercent, err := cpuKit.GetUsagePercent()
		if err != nil {
			mStats.CpuUsagePercentError = err
		} else {
			mStats.CpuUsagePercent = mathKit.Round(usagePercent, 2)
		}
	}()

	// (2) disk
	wg.Add(1)
	go func() {
		defer wg.Done()

		stats, err := diskKit.GetDiskUsageStats()
		if err != nil {
			mStats.DiskUsagePercentError = err
		} else {
			mStats.DiskPath = stats.Path
			mStats.DiskUsagePercent = mathKit.Round(stats.UsedPercent, 2)
		}
	}()

	// (3) others
	wg.Add(1)
	go func() {
		defer wg.Done()

		count, err := processKit.GetProcessCount()
		if err != nil {
			mStats.ProcessCountError = err
		} else {
			mStats.ProcessCount = count
		}

		count1, err := processKit.GetProcessThreadCount()
		if err != nil {
			mStats.ProcessThreadCountError = err
		} else {
			mStats.ProcessThreadCount = count1
		}

		stats, err := memoryKit.GetMachineMemoryStats()
		if err != nil {
			mStats.MemoryStatsError = err
		} else {
			mStats.Total = dataSizeKit.ToReadableIecString(float64(stats.Total))
			mStats.Available = dataSizeKit.ToReadableIecString(float64(stats.Available))
			mStats.Used = dataSizeKit.ToReadableIecString(float64(stats.Used))
			mStats.UsedPercent = mathKit.Round(stats.UsedPercent, 2)
			mStats.Free = dataSizeKit.ToReadableIecString(float64(stats.Free))
		}

		// ulimit -u
		if tmp, err := osKit.GetMaxProcessThreadCountByUser(); err != nil {
			mStats.MaxProcessThreadCountByUserError = err.Error()
		} else {
			mStats.MaxProcessThreadCountByUser = tmp
		}
		// kernel.pid_max
		if tmp, err := osKit.GetPidMax(); err != nil {
			mStats.PidMaxError = err.Error()
		} else {
			mStats.PidMax = tmp
		}
		// kernel.threads-max
		if tmp, err := osKit.GetThreadsMax(); err != nil {
			mStats.ThreadsMaxError = err.Error()
		} else {
			mStats.ThreadsMax = tmp
		}
		// vm.max_map_count
		if tmp, err := osKit.GetMaxMapCount(); err != nil {
			mStats.MaxMapCountError = err.Error()
		} else {
			mStats.MaxMapCount = tmp
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
	logger.Infof("stats:\n%s", json)
}
