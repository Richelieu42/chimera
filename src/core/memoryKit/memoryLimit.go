package memoryKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"math"
	"runtime/debug"
)

// SetSoftMemoryLimit 软内存限制(soft memory limit)
/*
PS:
(1) This limit will be respected even if GOGC=off (or, if SetGCPercent(-1) is executed).
(2) 默认memory limit是math.MaxInt64（即第一次调用debug.SetMemoryLimit()的返回值为math.MaxInt64）.

@param limit 单位: B(byte；字节)
@return 之前的内存限制
*/
func SetSoftMemoryLimit(limit uint64) (int64, error) {
	if limit < 512*dataSizeKit.MiB {
		return 0, errorKit.New("limit(%s) is too small", dataSizeKit.ToReadableIECString(limit))
	}
	if limit > math.MaxInt64 {
		return 0, errorKit.New("limit(%s) is too large", dataSizeKit.ToReadableIECString(limit))
	}

	return debug.SetMemoryLimit(int64(limit)), nil
}
