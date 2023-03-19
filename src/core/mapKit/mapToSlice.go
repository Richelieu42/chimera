package mapKit

import "github.com/samber/lo"

// MapToSlice map实例 => slice实例
/*
@param m		可以为nil
@param iteratee (1) 不能为nil，否则会导致panic: runtime error: invalid memory address or nil pointer dereference
					（当m==nil，iteratee可以为nil，但这样无意义）
				(2) 传参: 一对键值，返回值: slice实例中的一个元素
@return 必定不为nil（保底为len==0的slice实例）

e.g.
	s := mapKit.MapToSlice[string, string, string](map[string]string{"1": "a"}, func(key string, value string) string {
			return key + value
		})
	fmt.Println(s) // [1a]
*/
func MapToSlice[K comparable, V any, R any](m map[K]V, iteratee func(key K, value V) R) []R {
	return lo.MapToSlice(m, iteratee)
}
