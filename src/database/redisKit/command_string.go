package redisKit

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

// Set 设置指定key的值（string类型）.
/*
@param key 			可以为""
@param value 		支持的类型: string、[]byte、int、float64、bool(true: "1"; false: "0")...
					不支持的类型（会返回error）: map、自定义结构体...
@param expiration 	e.g.	120*time.Second			120s后过期
					 		0 						持久化的键（即TTL为-1），无论：键是否存在、存在的键是否有超时时间
					 		redis.KeepTTL(即-1) 	保持已经存在的TTL（需要确保Redis版本 >= 6.0，否则会返回error: ERR syntax error）
*/
func (client *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	reply, err := client.goRedisClient.Set(ctx, key, value, expiration).Result()
	if err != nil {
		return false, err
	}
	return reply == "OK", nil
}

// SetNX 只有在 key 不存在时设置 key 的值
/*
@return 第一个返回值代表: 是否设置成功
*/
func (client *Client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return client.goRedisClient.SetNX(ctx, key, value, expiration).Result()
}

// Get
/*
PS:
(1) 如果对应value的类型不为string，会返回error: WRONGTYPE Operation against a key holding the wrong kind of value

e.g.	当前db中不存在传参key
=>	("", redis.Nil)
*/
func (client *Client) Get(ctx context.Context, key string) (string, error) {
	return client.goRedisClient.Get(ctx, key).Result()
}

// GetWithoutRedisNil 对 Get 进行封装（特殊处理）: 当前db中不存在传参key时，返回 ("", nil).
/*
PS：
(1) 如果当前db中不存在传参key，将返回 ("", nil)
(2) 如果不关心key是否存在，只关心值，可以调用此方法
(3) 如果对应value的类型不为string，会返回error:	WRONGTYPE Operation against a key holding the wrong kind of value
*/
func (client *Client) GetWithoutRedisNil(ctx context.Context, key string) (string, error) {
	str, err := client.Get(ctx, key)
	if err != nil {
		if err != redis.Nil {
			return "", err
		}
		// err == redis.Nil: 当前db中不存在 传参key
		return "", nil
	}
	return str, nil
}
