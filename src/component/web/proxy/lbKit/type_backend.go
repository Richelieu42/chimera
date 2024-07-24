package lbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"net/http/httputil"
)

type Backend struct {
	mutexKit.Mutex

	// Alive 节点是否可用？
	Alive        bool
	ReverseProxy *httputil.ReverseProxy
}

func (be *Backend) SetAlive(alive bool) {
	be.LockFunc(func() {
		be.Alive = alive
	})
}

func (be *Backend) IsAlive() (alive bool) {
	be.LockFunc(func() {
		alive = be.Alive
	})
	return
}
