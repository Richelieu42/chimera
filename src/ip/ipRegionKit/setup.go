package ipRegionKit

import (
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
)

var NotSetupError = errorKit.Newf("haven’t been set up correctly")

// 缓存整个xdb数据的情况下，searcher对象可以安全用于并发
var searcher *xdb.Searcher

func MustSetUp(xdbPath string) {
	err := SetUp(xdbPath)
	if err != nil {
		zapKit.Fatalf("fail to set up, error: %s", err.Error())
	}
}

func SetUp(xdbPath string) (err error) {
	defer func() {
		if err != nil {
			searcher = nil
		}
	}()

	searcher, err = loadXdbFile(xdbPath)
	return err
}

// loadXdbFile
/*
@param path xdb文件的路径
*/
func loadXdbFile(xdbPath string) (*xdb.Searcher, error) {
	if err := fileKit.AssertExistAndIsFile(xdbPath); err != nil {
		return nil, err
	}

	// 缓存整个xdb数据
	cBuff, err := xdb.LoadContentFromFile(xdbPath)
	if err != nil {
		return nil, err
	}
	return xdb.NewWithBuffer(cBuff)
}
