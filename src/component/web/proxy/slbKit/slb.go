package slbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/atomic/atomicKit"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
)

func NewLoadBalancer() (lb *LoadBalancer) {
	// TODO:
	lb = &LoadBalancer{
		RWMutex: mutexKit.RWMutex{},

		backends: nil,

		current: atomicKit.NewInt32(-1),
	}
	return
}
