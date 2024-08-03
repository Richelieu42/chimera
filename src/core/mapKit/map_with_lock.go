package mapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
)

type (
	MapWithLock[K comparable, V any] struct {
		mutexKit.RWMutex

		// Map 并发不安全的
		Map map[K]V
	}
)

func (m *MapWithLock[K, V]) Size() (size int) {
	if m == nil {
		return
	}

	/* 读锁 */
	m.RLockFunc(func() {
		size = len(m.Map)
	})
	return
}

func NewMapWithLock[K comparable, V any]() *MapWithLock[K, V] {
	return &MapWithLock[K, V]{
		Map: make(map[K]V),
	}
}
