package slbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
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
	current *atomic.Int64

	// retry
	retry int16
	// retryInterval
	retryInterval time.Duration

	//// attempt
	//attempt int16

	// c 用于定期健康检查
	c *cron.Cron

	status Status
}

func (lb *LoadBalancer) AddBackend(backend *Backend) (err error) {
	if backend == nil {
		return
	}

	/* 写锁 */
	lb.LockFunc(func() {
		switch lb.status {
		case StatusInitialized, StatusStarted:
			// 允许添加后端服务
			lb.backends = append(lb.backends, backend)
		case StatusDisposed:
			err = AlreadyDisposedError
		default:
			err = errorKit.Newf("invalid status: %s", lb.status)
		}
	})
	return
}

// NextIndex 返回下一个下标.
/*
PS:
(1) 此方法无需加锁;
(2) 需要自行对返回值先进行 其余操作(%) 才能继续使用.
*/
func (lb *LoadBalancer) NextIndex() int64 {
	return lb.current.Inc()
}

// HealthCheck 对目前所有的后端服务，进行1次健康检查.
func (lb *LoadBalancer) HealthCheck() {
	/* 读锁 */
	lb.RLockFunc(func() {
		// 只有 StatusStarted 情况下，才会进行健康检查
		switch lb.status {
		case StatusInitialized:
			return
		case StatusStarted:
			// do nothing
		case StatusDisposed:
			return
		default:
			return
		}

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

// Start 手动启动.
func (lb *LoadBalancer) Start() error {
	/* 写锁 */
	lb.Lock()
	defer lb.Unlock()

	switch lb.status {
	case StatusInitialized:
		lb.status = StatusStarted
	case StatusStarted:
		return AlreadyStartedError
	case StatusDisposed:
		return AlreadyDisposedError
	default:
		return errorKit.Newf("invalid status: %s", lb.status)
	}

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
	lb.LockFunc(func() {
		cronKit.StopCron(lb.c)

		lb.backends = nil
		lb.c = nil
		lb.status = StatusDisposed
	})
}

func (lb *LoadBalancer) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	// getAccessBackend 从start下标开始获取可用服务（向后寻找）.
	/*
		PS: 调用此函数前，需要先获取读锁.
	*/
	getAccessBackend := func(start, limit, length int64) (*Backend, int64) {
		for i := start; i < limit; i++ {
			idx := i % length
			be := lb.backends[idx]
			if !be.IsAlive() {
				continue
			}
			// (1) 找到可用服务
			return be, i
		}
		// (2) 找不到可用服务
		return nil, 0
	}

	lb.RLock()
	defer lb.RUnlock()

	length := int64(len(lb.backends))
	start := lb.NextIndex() % length
	limit := start + length

	for i := start; i < limit; i++ {
		be, idx := getAccessBackend(i, limit, length)
		if be == nil {
			return NoAccessBackendError
		}
		i = idx
		// TODO: be.HandleRequest

	}

	//length := int32(len(lb.backends))
	//startIndex := lb.NextIndex() % length
	//limit := startIndex + length
	//for i := startIndex; i < limit; i++ {
	//
	//}

	return nil
}
