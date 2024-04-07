package timeKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"time"
)

// Interval
/*
参考:
Golang正确停止Ticker https://blog.csdn.net/weixin_40098405/article/details/111517279

!!!: 使用 time.Ticker 时要注意: Stop会停止Ticker，停止后，Ticker不会再被发送，但是Stop不会关闭通道，防止读取通道发生错误.
*/
type Interval struct {
	mutexKit.RWMutex

	stopped bool

	ticker *time.Ticker

	// closeCh 关闭通道.
	closeCh chan struct{}
}

// Stop
/*
PS:
(1) 可以多次调用，不会panic，但这样没意义（只有第一次调用才有意义）;
(2) 如果有任务正在执行，会等它先执行完.
*/
func (i *Interval) Stop() {
	if i == nil || i.stopped {
		return
	}

	/* 写锁 */
	i.LockFunc(func() {
		if i.stopped {
			return
		}

		i.stopped = true
		i.ticker.Stop()
		i.closeCh <- struct{}{}
	})
}

// SetInterval 效果类似于JavaScript中的 window.setInterval().
/*
@param task		(1) 不能为nil
				(2) 传参t为执行任务时的 time.Time
@param duration 必须>0
*/
func SetInterval(task func(t time.Time), duration time.Duration) *Interval {
	i := &Interval{
		RWMutex: mutexKit.RWMutex{},
		stopped: false,
		ticker:  time.NewTicker(duration),
		closeCh: make(chan struct{}),
	}

	go func(i *Interval) {
		//// test
		//defer func() {
		//	logrus.Info("goroutine ends")
		//}()

		defer i.ticker.Stop()

		for {
			select {
			case t := <-i.ticker.C:
				/* 读锁 */
				i.RLockFunc(func() {
					if !i.stopped {
						task(t)
					}
				})
			case <-i.closeCh:
				return
			}
		}
	}(i)
	return i
}

// SetInterval 效果类似于JavaScript中的同名函数.
//var SetInterval func(task func(t time.Time), duration time.Duration) *Interval = SetInterval

// ClearInterval 效果类似于JavaScript中的同名函数.
/*
@param i 可以为nil
*/
func ClearInterval(i *Interval) {
	i.Stop()
}
