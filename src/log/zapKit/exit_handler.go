package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
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
	/* 加锁 && 解锁 */
	mutex.LockFunc(func() {
		for _, handler := range handlers {
			if handlers == nil {
				continue
			}
			syncHandlers = append(syncHandlers, handler)
		}
	})
}

func RegisterParallelExitHandler(handlers ...func()) {
	/* 加锁 && 解锁 */
	mutex.LockFunc(func() {
		for _, handler := range handlers {
			if handlers == nil {
				continue
			}
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
		Info("No exit handler.")
		return
	}

	var wg sync.WaitGroup

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
	case <-time.After(timeout):
		Errorf("Fail to run all exit handlers within timeout(%s).", timeout)
	case <-endCh:
		Infof("Manager to run all exit handlers within timeout(%s).", timeout)
	}
}

// 参考: logrus.RegisterExitHandler
func runExitHandler(handler func()) {
	defer func() {
		if err := recover(); err != nil {
			Errorf("Recover from panic: %v", err)
		}
	}()

	handler()
}
