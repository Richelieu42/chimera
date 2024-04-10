package timeKit

import (
	"time"
)

var (
	Local *time.Location = time.Local

	UTC *time.Location = time.UTC

	GMT *time.Location
)

// ConvertLocation 时区转换.
/*
PS: 返回值是个副本，不会修改传参t.

@param loc 目标时区

e.g. UTC+8(CST) 转 UTC+0
	2022-05-05 14:33:40.562899 +0800 CST m=+0.001585418 => 2022-05-05 06:33:40.562899 +0000 UTC
*/
func ConvertLocation(t time.Time, loc *time.Location) time.Time {
	if loc == nil {
		loc = time.Local
	}
	return t.In(loc)
}

// LoadLocation
/*
LoadLocation的输入参数的取值，除了该函数的源代码中可看到的”UTC”、”Local”，其余的值其实是遵照“IANA Time Zone”的规则，可以解压$GOROOT/lib/time/zoneinfo.zip 这个文件打开查看。
在Asia这个目录，我看到了Chongqing，Hong_Kong，但没Beijing。在国外获取中国北京时间，要用”PRC”，当然”Asia/Chongqing”也是个方法
参考：https://blog.csdn.net/qq_26981997/article/details/53454606

@param name (a) 可以为""（将返回 time.UTC）
			(b) e.g. "Asia/Chongqing"
*/
var LoadLocation func(name string) (*time.Location, error) = time.LoadLocation
