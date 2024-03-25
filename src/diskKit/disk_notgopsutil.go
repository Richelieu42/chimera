//go:build !(386 || amd64 || arm || arm64)

package diskKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/conditionKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/osKit"
	"github.com/sirupsen/logrus"
)

func GetDiskUsageStats() (*DiskUsageStats, error) {
	path := conditionKit.TernaryOperator(osKit.IsWindows(), "C:", "/")
	return GetDiskUsageStatsByPath(path)
}

func GetDiskUsageStatsByPath(path string) (*DiskUsageStats, error) {
	return nil, errorKit.Newf("Currently not supported")
}

func PrintBasicDetails(logger *logrus.Logger) {

}
