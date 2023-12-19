package redisKit

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
)

// IsStreamSupported
/*
PS:
(1) 低版本Redis（<5）不支持Stream;
(2) Tendis 2.6.0不支持Stream.
*/
func (client *Client) IsStreamSupported(ctx context.Context) error {
	id := idKit.NewXid()
	stream := fmt.Sprintf("%s:%s:%s:%s", consts.ProjectName, "test", "redis-stream", id)

	defer func() {
		_, _ = client.Del(ctx, stream)
	}()

	cmd := client.universalClient.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: map[string]interface{}{
			"data": "test",
		},
	})
	_, err := cmd.Result()
	return err
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

// XGroupCreate [消费者] 创建消费者组.("xgroup", "create", stream, group, start)
/*
@Deprecated: Use XGroupCreateMkStream instead.

PS:
(1) 如果 stream 对应的key:	(a) 存在，do nothing;
							(b) 不存在，将返回error（ERR The XGROUP subcommand requires the key to exist. Note that for CREATE you may want to use the MKSTREAM option to create an empty stream automatically.）.
(2) 如果 group 已经存在，将返回error(BUSYGROUP Consumer Group name already exists).
*/
func (client *Client) XGroupCreate(ctx context.Context, stream, group, start string) error {
	resp, err := client.universalClient.XGroupCreate(ctx, stream, group, start).Result()
	if err != nil {
		return err
	}
	if !strKit.EqualsIgnoreCase(resp, "OK") {
		return errorKit.New("invalid resp(%s)", resp)
	}
	return nil
}

// XGroupCreateMkStream [消费者] 创建消费者组.("xgroup", "create", stream, group, start, "mkstream")
/*
PS:
(1) 如果 stream 对应的key:	(a) 存在，do nothing;
							(b) 不存在，将自动创建一个空的stream.
(2) 如果 group 已经存在，将返回error(BUSYGROUP Consumer Group name already exists).
(3) MKSTREAM是一个可选子命令，如果指定了它，那么在创建消费者组的时候，如果stream不存在，那么会自动创建一个空（长度为0）的stream.

@param group 	消费者组
@param start 	(1) "0": 从头开始消费
				(2) "$": 从末尾开始消费
*/
func (client *Client) XGroupCreateMkStream(ctx context.Context, stream, group, start string) error {
	resp, err := client.universalClient.XGroupCreateMkStream(ctx, stream, group, start).Result()
	if err != nil {
		return err
	}
	if !strKit.EqualsIgnoreCase(resp, "OK") {
		return errorKit.New("invalid resp(%s)", resp)
	}
	return nil
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
