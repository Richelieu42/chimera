// Package runtimeKit 主要是对如下包的封装："runtime"
package runtimeKit

import (
	"github.com/shirou/gopsutil/v3/host"
	"runtime"
)

func GetHostInfo() (*host.InfoStat, error) {
	return host.Info()
}

// GetGoRoot 环境变量GOROOT
func GetGoRoot() string {
	return runtime.GOROOT()
}

// GetGoVersion Golang的版本号
func GetGoVersion() string {
	return runtime.Version()
}
