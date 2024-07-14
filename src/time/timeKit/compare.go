package timeKit

import (
	"github.com/samber/lo"
	"time"
)

var (
	// Earliest 返回最小值
	Earliest func(times ...time.Time) time.Time = lo.Earliest

	// Latest 返回最大值
	Latest func(times ...time.Time) time.Time = lo.Latest
)

func EarliestBy[T any](collection []T, iteratee func(item T) time.Time) T {
	return lo.EarliestBy(collection, iteratee)
}

func LatestBy[T any](collection []T, iteratee func(item T) time.Time) T {
	return lo.LatestBy(collection, iteratee)
}

// Between 检查给定的时间是否处于某一时间区间内（左右都不包含！！！）
/*
参考:
【收藏】开发常用的 10 个通用函数 https://mp.weixin.qq.com/s/tvy9L-pb_8WFWAmA9u-bMg

e.g.
(cur, cur, cur.Add(time.Second)) 					=> false
(cur, cur.Add(-time.Second), cur.Add(time.Second)) 	=> true
*/
func Between(t, min, max time.Time) bool {
	return t.After(min) && t.Before(max)
}
