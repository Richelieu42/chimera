package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// XAdd 添加消息到末尾.
func (client *Client) XAdd(ctx context.Context, a *redis.XAddArgs) (string, error) {
	cmd := client.universalClient.XAdd(ctx, a)
	return cmd.Result()
}

// XDel 删除消息.
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
