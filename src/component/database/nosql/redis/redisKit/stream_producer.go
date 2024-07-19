package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// XAdd [生产者] 添加消息到末尾.
/*
PS:
(1) 传参stream 不存在，会自动创建;
(2) 传参stream 存在但类型不是 stream，会返回error: WRONGTYPE Operation against a key holding the wrong kind of value

语法: XADD key ID field value [field value ...]
key:			队列名称，如果不存在就创建
ID:				消息 id，我们使用 * 表示由 redis 生成，可以自定义，但是要自己保证递增性。
field value:	记录

@param a 	(1) 必需的字段: Stream、Values
				推荐的字段: MaxLen（防止消息堆积）、Approx（可选）
			(2) Stream字段对应: Redis中的key（stream类型）
			(3) 可选的ID字段，为 ""（默认） 则由Redis生成
@return 	id: 消息的id

e.g.
	_, err := impl.client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: topic,
		Values: map[string]interface{}{
			"tag":  tag,
			"data": data,
		},
	})
*/
func (client *Client) XAdd(ctx context.Context, a *redis.XAddArgs) (id string, err error) {
	cmd := client.universalClient.XAdd(ctx, a)
	id, err = cmd.Result()
	return
}
