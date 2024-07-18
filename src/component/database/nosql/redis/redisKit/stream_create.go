package redisKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

// XGroupCreate 创建消费者组（Consumer Group）.
/*
("xgroup", "create", stream, group, start)

!!!: 返回的err非nil的话，建议先进行 IsConsumerGroupNameAlreadyExistError 判断，返回true的话忽略该err.

PS:
(1) 如果 stream 对应的key:	(a) 存在，do nothing;
							(b) 不存在，将返回error（ERR The XGROUP subcommand requires the key to exist. Note that for CREATE you may want to use the MKSTREAM option to create an empty stream automatically.）.
(2) 如果 group 已经存在，将返回error(BUSYGROUP Consumer Group name already exists).
*/
func (client *Client) XGroupCreate(ctx context.Context, stream, group, start string) error {
	resp, err := client.universalClient.XGroupCreate(ctx, stream, group, start).Result()
	if err != nil {
		return errorKit.Wrapf(err, "fail with stream(%s), group(%s) and start(%s)", stream, group, start)
	}
	if !strKit.EqualsIgnoreCase(resp, "OK") {
		return errorKit.Newf("invalid resp(%s)", resp)
	}
	return nil
}

// XGroupCreateMkStream 创建消费者组（Consumer Group），如果stream不存在，那么会自动创建一个空（长度为0）的stream.
/*
("xgroup", "create", stream, group, start, "mkstream")

!!!: 返回的err非nil的话，建议先进行 IsConsumerGroupNameAlreadyExistError 判断，返回true的话忽略该error.

PS:
(1) 如果 stream 对应的key:	(a) 存在，do nothing;
							(b) 不存在，将自动创建一个空的stream.
(2) 如果 group 已经存在，将返回error(BUSYGROUP Consumer Group name already exists).
(3) MKSTREAM是一个可选子命令，如果指定了它，那么在创建消费者组的时候，如果stream不存在，那么会自动创建一个空（长度为0）的stream.

@param stream 	要创建的 消费者组的流 的名称
@param group 	要创建的 消费者组 的名称
@param start 	起始 ID，可以是具体的 ID 或特殊的 "$"
				e.g.
				(1) "0": 从头开始消费
				(2) "$": 从末尾开始消费（从流的最新消息开始）
*/
func (client *Client) XGroupCreateMkStream(ctx context.Context, stream, group, start string) error {
	resp, err := client.universalClient.XGroupCreateMkStream(ctx, stream, group, start).Result()
	if err != nil {
		return errorKit.Wrapf(err, "fail with stream(%s), group(%s) and start(%s)", stream, group, start)
	}
	if !strKit.EqualsIgnoreCase(resp, "OK") {
		return errorKit.Newf("invalid resp(%s)", resp)
	}
	return nil
}
