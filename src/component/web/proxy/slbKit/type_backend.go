package slbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"net/http/httputil"
	"net/url"
)

type Backend struct {
	mutexKit.Mutex

	// Alive 节点是否可用？
	Alive        bool
	URL          *url.URL
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
