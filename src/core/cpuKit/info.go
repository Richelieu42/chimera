package cpuKit

import (
	"github.com/klauspost/cpuid/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"runtime"
	"time"
)

// InVirtualMachine 是否在虚拟机中？
var InVirtualMachine func() bool = cpuid.CPU.VM

func GetVendorID() cpuid.Vendor {
	return cpuid.CPU.VendorID
}

// GetVendorString CPU供应商
/*
@return e.g."Apple"
*/
func GetVendorString() string {
	return cpuid.CPU.VendorString
}

// GetBrandName CPU品牌名称
/*
@return e.g."Apple M1 Pro"
*/
func GetBrandName() string {
	return cpuid.CPU.BrandName
}

func GetPhysicalCores() int {
	return cpuid.CPU.PhysicalCores
}

func GetThreadsPerCore() int {
	return cpuid.CPU.ThreadsPerCore
}

func GetLogicalCores() int {
	return cpuid.CPU.LogicalCores
}

// GetCpuNumber returns the number of logical CPUs usable by the current process.
var GetCpuNumber func() int = runtime.NumCPU

var GetFeatureSet func() []string = cpuid.CPU.FeatureSet

func GetFamily() int {
	return cpuid.CPU.Family
}

func GetModel() int {
	return cpuid.CPU.Model
}

// GetUsage CPU使用率
/*
e.g.
() =>
*/
func GetUsage() (float64, error) {
	s, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return s[0], nil
}
