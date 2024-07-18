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

// XAdd [生产者] 添加消息到末尾（如果指定的队列不存在，则创建一个队列）.
/*
语法: XADD key ID field value [field value ...]
key:			队列名称，如果不存在就创建
ID:				消息 id，我们使用 * 表示由 redis 生成，可以自定义，但是要自己保证递增性。
field value:	记录

@param a 	(1) 必需字段: Stream、Values
			(2) Stream字段对应: Redis中的key（stream类型）
			(3) 可选的ID字段，为 ""（默认） 则由Redis生成
@return 	id: 消息的id
*/
func (client *Client) XAdd(ctx context.Context, a *redis.XAddArgs) (id string, err error) {
	cmd := client.universalClient.XAdd(ctx, a)
	id, err = cmd.Result()
	return
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

// XReadGroup [消费者] 读取消费者组中的消息.("xreadgroup", "group")
/*
XReadGroupArgs结构体:
	Group 		消费组名
	Consumer	消费者名
	Count		读取数量
	Block		阻塞时间
	Streams		要读取的所有Stream（!!!: (1) 数量应当>=2; (2) 最后一个元素应该是 ">"）
*/
func (client *Client) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) ([]redis.XStream, error) {
	cmd := client.universalClient.XReadGroup(ctx, a)
	return cmd.Result()
}

// XAck [消费者] 将消息标记为"已处理".
/*
PS: 并不会删除对应消息.
*/
func (client *Client) XAck(ctx context.Context, stream, group string, ids ...string) (int64, error) {
	cmd := client.universalClient.XAck(ctx, stream, group, ids...)
	return cmd.Result()
}
