package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"time"
)

func main() {
	timer := timeKit.SetTimeout(func() {
		fmt.Println(timeKit.FormatCurrentTime())
	}, time.Second*5)
	timeKit.ClearTimeout(timer)

	fmt.Println(timeKit.FormatCurrentTime() + "--------------------------")
	time.Sleep(time.Second * 11)
	fmt.Println(timeKit.FormatCurrentTime() + "--------------------------")
}
