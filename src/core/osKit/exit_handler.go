package osKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"sync"
	"time"
)

var (
	mutex = new(mutexKit.Mutex)

	timeout = time.Second * 30

	handlers []func()

	parallelHandlers []func()
)

func RegisterExitHandler(handler func()) {
	if handler == nil {
		return
	}

	mutex.LockFunc(func() {
		handlers = append(handlers, handler)
	})
}

func RegisterParallelExitHandler(handler func()) {
	if handler == nil {
		return
	}

	mutex.LockFunc(func() {
		parallelHandlers = append(parallelHandlers, handler)
	})
}

func SetExitTimeout(d time.Duration) {
	if d <= 0 {
		return
	}

	timeout = d
}

func RunExitHandlers() {
	if len(handlers) == 0 && len(parallelHandlers) == 0 {
		return
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	/* 串行 */
	wg.Add(1)
	go func() {
		defer wg.Done()

		for _, handler := range handlers {
			runExitHandler(handler)
		}
	}()

	/* 并行 */
	wg.Add(1)
	go func() {
		defer wg.Done()

		var wg1 sync.WaitGroup
		for _, handler := range parallelHandlers {
			wg1.Add(1)
			go func() {
				defer wg1.Done()
				runExitHandler(handler)
			}()
		}
		wg1.Done()
	}()

	endCh := make(chan struct{})
	go func() {
		wg.Wait()
		endCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		zapKit.Errorf("Fail to run all exit handlers within timeout(%s)", timeout)
	case <-endCh:
		zapKit.Infof("Manager to run all exit handlers within timeout(%s)", timeout)
	}
}

// 参考: logrus.RegisterExitHandler
func runExitHandler(handler func()) {
	defer func() {
		if err := recover(); err != nil {
			zapKit.Errorf("recover from panic: %v", err)
		}
	}()

	handler()
}
