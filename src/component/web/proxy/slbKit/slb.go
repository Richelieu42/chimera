package slbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/atomic/atomicKit"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
)

// NewLoadBalancer
/*
PS: 返回的*LoadBalancer实例，需要手动调用 Start 以启动.
*/
func NewLoadBalancer(logger *zap.Logger) (lb *LoadBalancer) {
	if logger == nil {
		logger = zapKit.NewLogger(nil)
	}

	lb = &LoadBalancer{
		RWMutex: &mutexKit.RWMutex{},

		logger:   logger,
		backends: nil,
		current:  atomicKit.NewInt64(-1),
		status:   StatusInitialized,
	}
	return
}
