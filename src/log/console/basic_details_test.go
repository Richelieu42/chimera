package console

import (
	"testing"
	"time"
)

func TestPrintBasicDetails(t *testing.T) {
	PrintBasicDetails()

	// 等一下异步输出（网络时间）
	time.Sleep(time.Second * 3)
}
