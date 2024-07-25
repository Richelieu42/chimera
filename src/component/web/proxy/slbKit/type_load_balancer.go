package slbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"go.uber.org/atomic"
)

type LoadBalancer struct {
	mutexKit.RWMutex

	backends []*Backend

	// current
	current *atomic.Uint64
}

// NextIndex
/*
PS:
(1) 此方法无需加锁;
(2) 需要自行对返回值先进行 其余操作(%) 才能继续使用.
*/
func (lb *LoadBalancer) NextIndex() uint64 {
	return lb.current.Inc()
}

func (lb *LoadBalancer) AddBackend(backend *Backend) {
	if backend == nil {
		return
	}

	/* 写锁 */
	lb.LockFunc(func() {
		lb.backends = append(lb.backends, backend)
	})
}
