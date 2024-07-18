package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
)

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
PS:
(1) 传参stream 不存在，
(2) 传参stream 存在但类型不是 stream，
(3) 传参group 不存在.

XReadGroupArgs结构体:
	Group 		消费组名
	Consumer	消费者名
	Count		读取数量
	Block		阻塞时间
	Streams		要读取的所有Stream（!!!: (1) 数量应当>=2; (2) 最后一个元素应该是 ">"）

e.g.
	entries, err := impl.client.XReadGroup(context.Background(), &redis.XReadGroupArgs{
		Consumer: consumer,
		Group:    group,
		Streams:  []string{stream, ">"},
		Count:    10,
		//Block:    0,
		//NoAck:    false,
	})
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
