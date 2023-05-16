package redisKit

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
)

// NewDistributedMutex 生成Redis分布式互斥锁.
/*
PS:
(1) 不可重入锁;
(2) 更多详见"Redis分布式锁（多语言）.docx";
(3) 写入Redis中的键，默认TTL为 8s.

@param name 建议以 "mutex:" 为前缀

e.g. 将TTL修改为30s
NewDistributedMutex("name", redsync.WithExpiry(time.Second * 30))
*/
func (client *Client) NewDistributedMutex(name string, options ...redsync.Option) *redsync.Mutex {
	pool := goredis.NewPool(client.core) // or, pool := redigo.NewPool(...)
	sync := redsync.New(pool)
	return sync.NewMutex(name, options...)
}
