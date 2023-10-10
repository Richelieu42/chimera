package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// XAdd [生产者] 添加消息到末尾（如果指定的队列不存在，则创建一个队列）.
func (client *Client) XAdd(ctx context.Context, a *redis.XAddArgs) (string, error) {
	cmd := client.universalClient.XAdd(ctx, a)
	return cmd.Result()
}

// XDel 删除消息.
func (client *Client) XDel(ctx context.Context, stream string, ids ...string) (int64, error) {
	cmd := client.universalClient.XDel(ctx, stream, ids...)
	return cmd.Result()
}

// XGroupCreate [消费者] 创建消费者组.
func (client *Client) XGroupCreate(ctx context.Context, stream, group, start string) (string, error) {
	cmd := client.universalClient.XGroupCreate(ctx, stream, group, start)
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

// XReadGroup [消费者] 读取消费组中的消息.
func (client *Client) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) ([]redis.XStream, error) {
	cmd := client.universalClient.XReadGroup(ctx, a)
	return cmd.Result()
}
