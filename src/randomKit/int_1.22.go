//go:build go1.22

package randomKit

import "math/rand/v2"

// Int 生成随机int
/*
	PS:
	(0) 参考: github.com/duke-git/lancet/v2/random RandInt
	(1) 如果min == max，将返回 min;
	(2) 如果min > max，将交换两者的值.

	@param min 可以 < 0
	@param max 可以 < 0
	@return 范围: [min, max)
*/
func Int(min, max int) int {
	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return rand.IntN(max-min) + min
}
