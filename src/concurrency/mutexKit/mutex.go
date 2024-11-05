package mutexKit

import (
	"github.com/gogf/gf/v2/os/gmutex"
)

//type (
//	// Mutex 互斥锁.
//	/*
//		PS: 也可以让 结构体 继承 Mutex.
//	*/
//	Mutex struct {
//		gmutex.Mutex
//	}
//
//	// RWMutex 读写锁.
//	/*
//		PS: 也可以让 结构体 继承 RWMutex.
//	*/
//	RWMutex struct {
//		gmutex.RWMutex
//	}
//)

// NewMutex 互斥锁.
func NewMutex() *gmutex.Mutex {
	return &gmutex.Mutex{}
}

// NewRWMutex 读写锁.
func NewRWMutex() *gmutex.RWMutex {
	return &gmutex.RWMutex{}
}
