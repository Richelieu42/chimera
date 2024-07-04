package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"testing"
	"time"
)

func TestRegisterExitHandler(t *testing.T) {
	/* 串行（同步） */
	RegisterExitHandler(func() {
		time.Sleep(time.Second * 1)
		console.Info("[sync] 0")
	}, func() {
		time.Sleep(time.Second * 2)
		console.Info("[sync] 1")
	})

	/* 并行（异步） */
	RegisterParallelExitHandler(func() {
		time.Sleep(time.Second * 3)
		console.Info("[async] a")
	}, func() {
		time.Sleep(time.Second * 3)
		console.Info("[async] b")
	}, func() {
		time.Sleep(time.Second * 3)
		console.Info("[async] c")
	})

	console.Info("---")
	time.Sleep(time.Second * 1)
	console.Info("===")

	Exit(1)
}
