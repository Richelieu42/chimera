package mapKit

import (
	"github.com/richelieu42/go-scales/src/core/intKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
)

// Contains 判断 map实例 中是否存在 指定的key.
/*
@param m 可以为nil（此时返回值固定为false）

e.g.
(map[string]interface{}(nil), "1") => false
*/
func Contains[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// ContainKeys 判断 map实例 中是否存在 所有指定的key
/*
@param keys 可以一个key都不传，此时将固定返回true
*/
func ContainKeys[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, key := range keys {
		if !Contains(m, key) {
			return false
		}
	}
	return true
}

// GetKeySlice 获取map实例中的所有key
/*
@param m 	如果为 nil 或 空的map实例，将返回nil
@return 	非nil的slice实例（len >= 0）
*/
func GetKeySlice[K comparable, V any](m map[K]V) []K {
	s := make([]K, 0, len(m))

	for key := range m {
		s = append(s, key)
	}
	return s
}

// Remove
/*
PS: 可能会修改传参m（移除的话），因为它是map类型.

@return 被移除出map的条目的值（存在的话） + 传参m是否包含传参key
*/
func Remove[K comparable, V any](m map[K]V, key K) (V, bool) {
	value, exist := m[key]
	if exist {
		// 存在的话，移除对应条目
		delete(m, key)
	}
	return value, exist
}

// Set 设置值（或更新值）
/*
@param m 不能为nil（否则会导致 panic: assignment to entry in nil map）
*/
func Set[K comparable, V any](m map[K]V, key K, value V) {
	m[key] = value
}

func GetString[K comparable, V any](m map[K]V, key K) (string, error) {
	obj, exist := m[key]
	if !exist {
		// 不存在对应key的情况下，返回零值
		return "", nil
	}

	return strKit.ToStringE(obj)
}

func GetInt[K comparable, V any](m map[K]V, key K) (int, error) {
	obj, exist := m[key]
	if !exist {
		// 不存在对应key的情况下，返回零值
		return 0, nil
	}
	return intKit.ToIntE(obj)
}
