package slbKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/proxy/forwardKit"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Backend struct {
	mutexKit.Mutex

	// Alive 节点是否可用？
	Alive        bool
	URL          *url.URL
	ReverseProxy *httputil.ReverseProxy
}

func (be *Backend) Enable() {
	/* 锁 */
	be.LockFunc(func() {
		be.Alive = true
	})
}

func (be *Backend) Disable() {
	/* 锁 */
	be.LockFunc(func() {
		be.Alive = false
	})
}

func (be *Backend) IsAlive() (alive bool) {
	/* 锁 */
	be.LockFunc(func() {
		alive = be.Alive
	})
	return
}

// HealthCheck 健康检查（此方法会修改 Alive 字段）.
/*
@return 后端服务是否可用？
*/
func (be *Backend) HealthCheck() {
	timeout := 3 * time.Second

	conn, err := netKit.DialTimeout("tcp", be.URL.Host, timeout)
	if err != nil {
		be.Disable()
		return
	}
	_ = conn.Close()
	be.Enable()
}

func (be *Backend) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	return forwardKit.ForwardByReverseProxy(w, r, be.ReverseProxy)
}
