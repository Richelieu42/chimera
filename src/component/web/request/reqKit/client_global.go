package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
)

var (
	// globalClient 全局的客户端
	globalClient = NewClient()

	globalMutex = new(mutexKit.RWMutex)
)

func ReplaceGlobalClient(client *req.Client) {
	if client == nil {
		return
	}

	/* 写锁 */
	globalMutex.LockFunc(func() {
		globalClient = client
	})
}

func GetGlobalClient() (client *req.Client) {
	/* 读锁 */
	globalMutex.RLockFunc(func() {
		client = globalClient
	})
	return
}
