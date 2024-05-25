package netKit

import (
	"github.com/richelieu-yang/chimera/v3/src/reflectKit"
	"reflect"
	"strconv"
	"time"
)

const (
	DefaultHttpPort  = 80
	DefaultHttpsPort = 443

	// MaxPort 65535 == 0xFFFF
	MaxPort = 0xFFFF
)

// IsValidPort
/*
参考:
(1) Java，hutool中的NetUtil.isValidPort()
(2) Linux端口分配: https://blog.csdn.net/zh2508/article/details/104888743

0 			不使用
1–1023 		系统保留,只能由root用户使用
1024—4999 	由客户端程序自由分配
5000—65535 	由服务器端程序自由分配（65535 = 2 ^ 16 - 1）

@param obj 	(1) 支持的类型: reflect.Value、int、uint、string...
			(2) 可以为nil
*/
func IsValidPort(obj interface{}) bool {
	if obj == nil {
		return false
	}

	if v, ok := obj.(reflect.Value); ok {
		return isReflectValueValidPort(v)
	}
	return isReflectValueValidPort(reflectKit.ValueOf(obj))
}

func isReflectValueValidPort(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		i, err := strconv.Atoi(v.String())
		if err != nil {
			return false
		}
		return i > 0 && i <= MaxPort
	default:
		if v.CanInt() {
			i := v.Int()
			return i > 0 && i <= MaxPort
		} else if v.CanUint() {
			i := v.Uint()
			return i > 0 && i <= MaxPort
		}
		return false
	}
}

// IsLocalPortAvailable 本地端口是否可用（即未被占用）？
// Deprecated: 某些绑定非127.0.0.1的端口无法被检测到.
/*
PS: 会优先判断端口是否有效（valid）.

参考:
Java，hutool中的NetUtil.isUsableLocalPort()
golang端口占用检测的使用	https://wenku.baidu.com/view/25716f5b01768e9951e79b89680203d8ce2f6af5.html
*/
func IsLocalPortAvailable(port int) bool {
	if !IsValidPort(int64(port)) {
		return false
	}

	// 能连通就说明端口被占用了
	addr := JoinToHost("127.0.0.1", port)

	return CanDial(addr, time.Second*3) != nil
}
