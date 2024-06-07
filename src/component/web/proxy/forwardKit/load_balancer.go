package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"net/http"
)

type (
	LoadBalancer struct {
		mutexKit.RWMutex

		backends []*Backend
	}
)

func (lb *LoadBalancer) AddBackend(backend *Backend) {
	if backend == nil {
		return
	}

	/* 写锁 */
	lb.LockFunc(func() {
		lb.backends = append(lb.backends, backend)
	})
}

// Handle 负载均衡http请求.
func (lb *LoadBalancer) Handle(w http.ResponseWriter, t *http.Request) {

}
