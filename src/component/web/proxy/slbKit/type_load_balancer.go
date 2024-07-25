package slbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"go.uber.org/atomic"
)

type LoadBalancer struct {
	mutexKit.RWMutex

	backends []*Backend

	// current 当前的下标
	current *atomic.Int64
}

// NextIndex
/*
PS:
(1) 此方法无需加锁;
(2) 需要自行对返回值先进行 其余操作(%) 才能继续使用.
*/
func (lb *LoadBalancer) NextIndex() int64 {
	return lb.current.Inc()
}

// UpdateIndex
/*
PS: 此方法无需加锁.
*/
func (lb *LoadBalancer) UpdateIndex(i int64) {
	lb.current.Store(i)
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

// GetNextPeer
/*
@return 可能为nil（即此时无可用后端服务）
*/
func (lb *LoadBalancer) GetNextPeer() (backend *Backend) {
	/* 读锁 Richelieu: 此处使用读锁而非写锁，以提高并发效率 */
	lb.RLockFunc(func() {
		length := int64(len(lb.backends))
		if length == 0 {
			// 直接没服务
			return
		}

		next := lb.NextIndex() % int64(len(lb.backends))
		limit := next + length
		for i := next; i < limit; i++ {
			idx := i % length
			tmp := lb.backends[idx]
			if !tmp.IsAlive() {
				continue
			}
			// 成功获取到可用服务
			backend = tmp
			lb.current.Store(idx)
			return
		}
	})
	return
}
