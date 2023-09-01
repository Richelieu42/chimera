package redisKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v2/src/core/intKit"
)

// NewUniqueId 通过Redis生成唯一id.
/*
PS: 返回值可以用作 分布式唯一id.
*/
func (client *Client) NewUniqueId(ctx context.Context, key string) (string, error) {
	i, err := client.Incr(ctx, key)
	if err != nil {
		return "", err
	}
	return intKit.Int64ToString(i), nil
}
