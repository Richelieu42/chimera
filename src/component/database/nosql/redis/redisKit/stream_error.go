package redisKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

// IsConsumerGroupNameAlreadyExistError 适用场景:
/*
TODO: 涉及 github.com/redis/go-redis/v9 源码， 后续看有没有好的解决方法.

PS: 与 XGroupCreateMkStream 搭配使用.
*/
func IsConsumerGroupNameAlreadyExistError(err error) bool {
	if err == nil {
		return false
	}

	err = errorKit.Cause(err)
	return strKit.ContainsIgnoreCase(err.Error(), "BUSYGROUP Consumer Group name already exists")
}

// IsNoStreamOrNoGroupError 适用场景: Consumer读取消息时返回error.
/*
@return true: 错误是由 "指定stream不存在" 或 "指定stream存在，但指定group不存在" 导致的.
*/
func IsNoStreamOrNoGroupError(err error) bool {
	if err == nil {
		return false
	}

	err = errorKit.Cause(err)
	return strKit.ContainsIgnoreCase(err.Error(), "NOGROUP No such key") && strKit.ContainsIgnoreCase(err.Error(), "or consumer group")
}
