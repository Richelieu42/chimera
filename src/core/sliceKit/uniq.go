package sliceKit

import "github.com/samber/lo"

// Uniq 去重.
/*
PS: 不会修改传参s.

@param s 	可以为nil
@return		可能为nil

e.g.
	s := sliceKit.Uniq([]interface{}{0, 1, 2, 0, "1", "2", "1"})
	fmt.Println(s)	// [0 1 2 1 2]（前3个为int类型，后2个为string类型）
*/
func Uniq[T comparable](s []T) []T {
	if s == nil {
		return nil
	}
	return lo.Uniq(s)
}

// UniqBy
/*
PS: 不会修改传参s.

@param s 		可以为nil
@param iteratee 不能为nil
@return			必定不为nil（保底为空的slice实例）

e.g.
	s0 := []int{0, 1, 2, 3, 4, 5}
	s1 := sliceKit.UniqBy[int, int](s0, func(i int) int {
		return i % 3
	})
	fmt.Println(s0)	// [0 1 2 3 4 5]
	fmt.Println(s1)	// [0 1 2]
*/
func UniqBy[T any, U comparable](collection []T, iteratee func(item T) U) []T {
	return lo.UniqBy(collection, iteratee)
}
