package mapKit

import (
	"github.com/mohae/deepcopy"
	"github.com/richelieu42/go-scales/src/core/errorKit"
)

// Clone 浅拷贝.
func Clone[K comparable, V any](m map[K]V) map[K]V {
	if m == nil {
		return nil
	}

	dolly := make(map[K]V)
	for k, v := range m {
		dolly[k] = v
	}
	return dolly
}

// DeepClone 深拷贝.
/*
参考:
「Go工具箱」推荐一个非常简单的深拷贝工具：deepcopy https://mp.weixin.qq.com/s/e3bL1i6WT-4MwK-SEpUa6Q
*/
func DeepClone[K comparable, V any](m map[K]V) (map[K]V, error) {
	obj := deepcopy.Copy(m)

	if dolly, ok := obj.(map[K]V); ok {
		return dolly, nil
	}
	return nil, errorKit.Simple("fail to deep clone")
}
