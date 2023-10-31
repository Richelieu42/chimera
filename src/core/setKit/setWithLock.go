package setKit

import "github.com/richelieu-yang/chimera/v2/src/mutexKit"

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type (
	SetWithLock[T comparable] struct {
		RWLock mutexKit.RWMutex

		// Set 并发不安全的
		Set mapset.Set[T]
	}
)

func NewSetWithLock[T comparable]() *SetWithLock[T] {
	return &SetWithLock[T]{
		RWLock: mutexKit.RWMutex{},
		Set:    NewSet[T](false),
	}
}
