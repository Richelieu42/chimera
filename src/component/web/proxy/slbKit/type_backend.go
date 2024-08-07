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

// Backend 后端节点.
type Backend struct {
	mutexKit.Mutex

	// alive 当前节点是否可用？
	alive        bool
	u            *url.URL
	reverseProxy *httputil.ReverseProxy
}

func (be *Backend) Enable() {
	/* 锁 */
	be.LockFunc(func() {
		be.alive = true
	})
}

func (be *Backend) Disable() {
	/* 锁 */
	be.LockFunc(func() {
		be.alive = false
	})
}

func (be *Backend) IsAlive() (alive bool) {
	/* 锁 */
	be.LockFunc(func() {
		alive = be.alive
	})
	return
}

// HealthCheck 健康检查（此方法会修改 alive 字段）.
/*
@return 后端服务是否可用？
*/
func (be *Backend) HealthCheck() {
	timeout := 3 * time.Second

	conn, err := netKit.DialTimeout("tcp", be.u.Host, timeout)
	if err != nil {
		be.Disable()
		return
	}
	_ = conn.Close()
	be.Enable()
}

func (be *Backend) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	return forwardKit.ForwardByReverseProxy(w, r, be.reverseProxy)
}
