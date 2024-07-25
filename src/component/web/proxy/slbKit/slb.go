package slbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/atomic/atomicKit"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
)

func NewLoadBalancer() *LoadBalancer {
	// TODO:
	return &LoadBalancer{
		RWMutex: mutexKit.RWMutex{},

		backends: nil,

		current: atomicKit.NewInt64(-1),
	}
}
