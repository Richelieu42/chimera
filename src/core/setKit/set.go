package setKit

import (
	mapset "github.com/deckarep/golang-set/v2"
)

// NewSet
/*
@param threadSafe	是否并发安全？
*/
func NewSet[T comparable](threadSafe bool, args ...T) mapset.Set[T] {
	if threadSafe {
		return mapset.NewSet(args...)
	}
	return mapset.NewThreadUnsafeSet(args...)
}

func NewSetFromMapKeys[T comparable, V any](threadSafe bool, val map[T]V) mapset.Set[T] {
	if threadSafe {
		return mapset.NewSetFromMapKeys(val)
	}
	return mapset.NewThreadUnsafeSetFromMapKeys(val)
}

func NewSetWithSize[T comparable](threadSafe bool, cardinality int) mapset.Set[T] {
	if threadSafe {
		return mapset.NewSetWithSize[T](cardinality)
	}
	return mapset.NewThreadUnsafeSetWithSize[T](cardinality)
}
