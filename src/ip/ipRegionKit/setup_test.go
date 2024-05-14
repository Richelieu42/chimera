package ipRegionKit

import (
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetRegion(t *testing.T) {
	wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		panic(err)
	}
	logrus.Infof("wd: [%s].", wd)

	// https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.xdb
	xdbPath := "_chimera-lib/ip2region.xdb"
	MustSetUp(xdbPath)

	//ip := "10.0.9.141"
	ip := "218.90.174.146"
	logrus.Info(GetRegion(ip))
}
