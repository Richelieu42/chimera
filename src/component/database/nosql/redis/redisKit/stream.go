package redisKit

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
)

// IsStreamSupported
/*
PS:
(1) 低版本Redis（<5）不支持Stream;
(2) Tendis 2.6.0不支持Stream;
(3) TongRDS（具体版本未知）支持Stream.
*/
func (client *Client) IsStreamSupported(ctx context.Context) error {
	id := idKit.NewXid()
	stream := fmt.Sprintf("%s:%s:%s:%s", consts.ProjectName, "test", "redis-stream", id)

	defer func() {
		_, _ = client.Del(ctx, stream)
	}()

	_, err := client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: map[string]interface{}{
			"data": "test",
		},
	})
	if err != nil {
		return errorKit.Wrapf(err, "redis stream isn't supported")
	}
	return nil
}

// XDel 删除Stream中的特定消息.
/*
@return (1) 删除成功: 返回(1, nil)
		(2) 删除失败: 返回(0, nil)（e.g. stream 和 id 对应的消息不存在）
*/
func (client *Client) XDel(ctx context.Context, stream string, ids ...string) (int64, error) {
	cmd := client.universalClient.XDel(ctx, stream, ids...)
	return cmd.Result()
}

func (client *Client) XRead(ctx context.Context, a *redis.XReadArgs) ([]redis.XStream, error) {
	cmd := client.universalClient.XRead(ctx, a)
	return cmd.Result()
}

func (client *Client) XReadStreams(ctx context.Context, streams ...string) ([]redis.XStream, error) {
	cmd := client.universalClient.XReadStreams(ctx, streams...)
	return cmd.Result()
}
