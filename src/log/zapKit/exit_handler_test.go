package zapKit

import (
	"testing"
	"time"
)

func TestRegisterExitHandler(t *testing.T) {
	/* 串行（同步） */
	RegisterExitHandler(func() {
		time.Sleep(time.Second * 1)
		Info("[sync] 0")
	}, func() {
		time.Sleep(time.Second * 2)
		Info("[sync] 1")
	})

	/* 并行（异步） */
	RegisterParallelExitHandler(func() {
		time.Sleep(time.Second * 3)
		Info("[async] a")
	}, func() {
		time.Sleep(time.Second * 3)
		Info("[async] b")
	}, func() {
		time.Sleep(time.Second * 3)
		Info("[async] c")
	})

	Info("---")
	time.Sleep(time.Second * 1)
	Info("===")

	Exit(1)
}
