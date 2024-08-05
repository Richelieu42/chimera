package slbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/atomic/atomicKit"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
)

// NewLoadBalancer TODO:
/*
@return 需要手动调用 Start 以启动.
*/
func NewLoadBalancer() (lb *LoadBalancer) {
	lb = &LoadBalancer{
		RWMutex:  &mutexKit.RWMutex{},
		backends: nil,
		current:  atomicKit.NewInt64(-1),
		status:   StatusInitialized,
	}
	return
}
