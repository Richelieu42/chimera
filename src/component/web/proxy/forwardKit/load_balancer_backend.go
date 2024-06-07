package forwardKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	"net/http/httputil"
)

type (
	Backend struct {
		mutexKit.RWMutex

		Id string

		Access bool

		ReverseProxy *httputil.ReverseProxy
	}
)

func (be *Backend) SetAccess(flag bool) {
	/* 写锁 */
	be.LockFunc(func() {
		be.Access = flag
	})
}

func (be *Backend) IsAccess() (access bool) {
	/* 读锁 */
	be.RLockFunc(func() {
		access = be.Access
	})
	return
}

func NewBackend(rp *httputil.ReverseProxy) (*Backend, error) {
	if err := interfaceKit.AssertNotNil(rp, "rp"); err != nil {
		return nil, err
	}

	return &Backend{
		Id:           idKit.NewXid(),
		Access:       true,
		ReverseProxy: rp,
	}, nil
}
