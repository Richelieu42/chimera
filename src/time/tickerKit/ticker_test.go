package tickerKit

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTicker(t *testing.T) {
	// 创建一个每秒触发一次的Ticker
	ticker := NewTicker(time.Second)

	/*
		使用for循环持续接收并处理Tick事件
		!!!: 缺陷：虽然调用了Stop方法，但是 ticker.C通道 并没有关闭，导致goroutine无法退出.
	*/
	go func() {
		defer fmt.Println("goroutine ends")

		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// 等待 5s 后停止Ticker
	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped.")

	time.Sleep(time.Second * 3)
}
