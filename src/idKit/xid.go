package idKit

import (
	"github.com/rs/xid"
	"time"
)

// NewXid Xid是一个全局唯一的ID生成器，它使用Mongo Object ID算法来生成全局唯一的ID.
/*
PS:
(1) 单进程内生成唯一id，推荐使用 xid.
(2) 一个进程（Process）内，生成的id不会重复.

@return (1) 长度固定为20
		(2) 组成: 小写字母（a-f）、数字（0-9）
e.g.
	"ckic7hfnl531vbl645n0"
	"ckth51co47mgs2kacmk0"
*/
func NewXid() string {
	return xid.New().String()
}

func NewXidWithTime(t time.Time) string {
	return xid.NewWithTime(t).String()
}
