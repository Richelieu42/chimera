package ipToRegionKit

import (
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"testing"
)

func TestGetRegion(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		console.Infof("working dir: [%s].", wd)
	}

	/*
		https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.xdb
		下载下来，放到 "_chimera-lib" 目录下.
	*/
	xdbPath := "_chimera-lib/ip2region.xdb"
	MustSetUp(xdbPath)

	//ip := "10.0.9.141"
	ip := "49.93.11.38"
	str, err := GetRegion(ip)
	if err != nil {
		panic(err)
	}
	console.Info(str)
}
