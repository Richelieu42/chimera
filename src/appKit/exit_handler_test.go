package appKit

import (
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"testing"
	"time"
)

func TestRegisterExitHandler(t *testing.T) {
	/* 串行（同步） */
	RegisterExitHandler(func() {
		time.Sleep(time.Second * 1)
		zapKit.Info("[sync] 0")
	}, func() {
		time.Sleep(time.Second * 2)
		zapKit.Info("[sync] 1")
	})

	/* 并行（异步） */
	RegisterParallelExitHandler(func() {
		time.Sleep(time.Second * 3)
		zapKit.Info("[async] a")
	}, func() {
		time.Sleep(time.Second * 3)
		zapKit.Info("[async] b")
	}, func() {
		time.Sleep(time.Second * 3)
		zapKit.Info("[async] c")
	})

	zapKit.Info("---")
	time.Sleep(time.Second * 1)
	zapKit.Info("===")

	Exit()
}
