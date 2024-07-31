package slbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"github.com/richelieu-yang/chimera/v3/src/cronKit"
	"github.com/robfig/cron/v3"
	"go.uber.org/atomic"
	"net/http"
	"sync"
	"time"
)

type LoadBalancer struct {
	*mutexKit.RWMutex

	backends []*Backend

	// current 当前的下标
	current *atomic.Int32

	// retry
	retry int16
	// retryInterval
	retryInterval time.Duration

	// attempt
	attempt int16

	// c 用于定期健康检查
	c *cron.Cron

	disposed bool
}

// NextIndex 返回下一个下标.
/*
PS:
(1) 此方法无需加锁;
(2) 需要自行对返回值先进行 其余操作(%) 才能继续使用.
*/
func (lb *LoadBalancer) NextIndex() int32 {
	return lb.current.Inc()
}

// UpdateIndex 更新下标.
/*
PS: 此方法无需加锁.
*/
func (lb *LoadBalancer) UpdateIndex(i int32) {
	if i < 0 {
		i = 0
	}
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
	/*
		读锁
		Richelieu: 此处使用读锁而非写锁，以提高并发效率.
	*/
	lb.RLockFunc(func() {
		length := int32(len(lb.backends))
		if length == 0 {
			// 直接没服务
			return
		}

		next := lb.NextIndex() % int32(len(lb.backends))
		limit := next + length
		for i := next; i < limit; i++ {
			idx := i % length
			tmp := lb.backends[idx]
			if !tmp.IsAlive() {
				continue
			}
			// 成功获取到可用服务
			backend = tmp
			lb.UpdateIndex(idx)
			return
		}
	})
	return
}

// HealthCheck 对目前所有的后端服务，进行1次健康检查.
func (lb *LoadBalancer) HealthCheck() {
	/* 读锁 */
	lb.RLockFunc(func() {
		var wg sync.WaitGroup
		for _, backend := range lb.backends {
			wg.Add(1)
			go func() {
				defer wg.Done()
				backend.HealthCheck()
			}()
		}
		wg.Wait()
	})
}

func (lb *LoadBalancer) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	// TODO:

	return nil
}

// Start 手动启动.
func (lb *LoadBalancer) Start() error {
	/* 写锁 */
	lb.Lock()
	defer lb.Unlock()

	/* 以10s为周期，对所有后端服务进行健康检查 */
	c, _, err := cronKit.NewCronWithTask("@every 10s", func() {
		lb.HealthCheck()
	})
	if err != nil {
		return err
	}
	c.Start() // 不阻塞
	lb.c = c
	return nil
}

// Dispose 手动中止.
func (lb *LoadBalancer) Dispose() {
	/* 写锁 */
	lb.Lock()
	defer lb.Unlock()

	defer func() {
		lb.backends = nil
		lb.c = nil
		lb.disposed = true
	}()
	cronKit.StopCron(lb.c)
}
