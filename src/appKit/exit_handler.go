package appKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"sync"
	"time"
)

var (
	mutex = new(mutexKit.Mutex)

	timeout = time.Second * 30

	syncHandlers []func()

	asyncHandlers []func()
)

func RegisterExitHandler(handlers ...func()) {
	handlers = sliceKit.RemoveZeroValues(handlers)
	if len(handlers) == 0 {
		return
	}

	/* 加锁 && 解锁 */
	mutex.LockFunc(func() {
		for _, handler := range handlers {
			syncHandlers = append(syncHandlers, handler)
		}
	})
}

func RegisterParallelExitHandler(handlers ...func()) {
	handlers = sliceKit.RemoveZeroValues(handlers)
	if len(handlers) == 0 {
		return
	}

	/* 加锁 && 解锁 */
	mutex.LockFunc(func() {
		for _, handler := range handlers {
			asyncHandlers = append(asyncHandlers, handler)
		}
	})
}

func SetExitTimeout(d time.Duration) {
	if d <= 0 {
		return
	}

	timeout = d
}

func RunExitHandlers() {
	if len(syncHandlers) == 0 && len(asyncHandlers) == 0 {
		return
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	/* 串行 */
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, handler := range syncHandlers {
			runExitHandler(handler)
		}
	}()

	/* 并行 */
	wg.Add(1)
	go func() {
		defer wg.Done()

		var wg1 sync.WaitGroup
		for _, handler := range asyncHandlers {
			wg1.Add(1)
			go func() {
				defer wg1.Done()
				runExitHandler(handler)
			}()
		}
		wg1.Wait()
	}()

	endCh := make(chan struct{})
	go func() {
		wg.Wait()
		endCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		zapKit.Errorf("Fail to run all exit syncHandlers within timeout(%s).", timeout)
	case <-endCh:
		zapKit.Infof("Manager to run all exit syncHandlers within timeout(%s).", timeout)
	}
}

// 参考: logrus.RegisterExitHandler
func runExitHandler(handler func()) {
	defer func() {
		if err := recover(); err != nil {
			zapKit.Errorf("Recover from panic: %v", err)
		}
	}()

	handler()
}
