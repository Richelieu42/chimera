package sliceKit

import (
	"github.com/samber/lo"
	"reflect"
)

// Get 根据下标获取slice中的元素
/*
Deprecated: 考虑性能的场景下，不建议直接调用此方法（此方法仅供展示传参规范）.

PS:
(1) 如果s == nil，会导致panic；（不管index为何值，即使为0）
(2) 如果s != nil && len(s) == 0，会导致panic；（不管index为何值，即使为0）
(3) 如果s != nil && len(s) > 0，index的取值范围: [0, length).
*/
func Get[T any](s []T, index int) T {
	return s[index]
}

// Append 向slice实例的"最后面"添加元素
/*
Deprecated: 考虑性能的场景下，不建议直接调用此方法（此方法仅供展示传参规范）.

PS:
(1) 传参s == nil的情况下，此时如果eles数量>=1，将返回1个非nil的slice实例，否则将返回nil.
(2) append()返回的是1个新的slice实例.

@param s 	可以为nil
@return 	可能为nil

e.g.
([]string(nil)) 					=> nil
([]string(nil), []string(nil)...) 	=> nil
([]string(nil), "0")				=> []string{"0"}
*/
func Append[T any](s []T, eles ...T) []T {
	return append(s, eles...)
}

// Merge 合并多个切片（不会去重）
/*
PS:
(1) 先要传参nil的话，必须要造型. e.g. []string(nil)
(2) 第1个传参可以为nil

@return 可能为nil

e.g.
	s := sliceKit.Merge[string](nil, []string{})
	fmt.Println(s)        // []
	fmt.Println(len(s))   // 0
	fmt.Println(s != nil) // true

e.g.1
	([]string(nil), []string{"1", "2"}) => [1 2]
	([]string{"1", "2"}, []string(nil)) => [1 2]

	([]string(nil))	=> nil
	([]string{}) 	=> []

	([]string(nil), []string{"a", "b"}, []string(nil), []string{"b", "c"}) => [a b b c]
*/
func Merge[T comparable](slices ...[]T) []T {
	var rst []T

	for _, s := range slices {
		if rst == nil {
			rst = s
		} else {
			rst = append(rst, s...)
		}
	}
	return rst
}

// Shuffle 随机打乱切片.
/*
PS: 会改变传参s的内容.

@param s 可以为nil

e.g.
[string](nil) => nil
[]string{"0", "1", "2", "3"} => [2 0 3 1]
*/
func Shuffle[T any](s []T) []T {
	return lo.Shuffle(s)
}

// Reverse 反转切片.
/*
Deprecated: This helper is mutable. This behavior might change in v2.0.0.

@param s (1)可以为nil (2)可能会改变传参s的内容.
@return 可能为nil

e.g.
(nil) 							=> nil
([]string{"0", "1", "2", "3"}) 	=> [3 2 1 0]
*/
func Reverse[T any](s []T) []T {
	return lo.Reverse(s)
}

// Swap 交换切片实例中两个元素的值
/*
PS:
（1）传参s不能为nil，会panic（下标越界）；
（2）此方法会修改传参s（虽然golang是值传递）；
（3）传参i、j：下标越界会导致panic.

@param s 不能为nil
@param i 第1个元素的下标（从0开始）
@param j 第2个元素的下标（从0开始）

e.g.
	s := []int{1, 0, 100}
	sliceKit.Swap(s, 0, 2)
	fmt.Println(s)			// [100 0 1]
*/
func Swap[T any](s []T, i, j int) {
	s[i], s[j] = s[j], s[i]
}

// GetFirstElement 主要用于: 从不定参数(...)中取第一个值（不存在则取默认值）
/*
Deprecated: 不建议在本项目内部调用，以防import cycle.

PS:
(1) 因为Golang不支持方法重载；
(2) T类型值可能为nil的情况，要注意防坑.

@param args 要么是: nil；要么是: 长度>=1的切片实例
*/
func GetFirstElement[T any](def T, args ...T) T {
	if len(args) > 0 {
		return args[0]
	}
	return def
}

// Compact 去除零值.
/*
@param s 可以为nil；不会修改传参s
@return 必定不为nil（保底为空的slice实例）

e.g.
	s := []string{"", "foo", "", "bar", ""}
	s1 := sliceKit.Compact[string](s)
	fmt.Println(s1) 	// [foo bar]
*/
func Compact[T comparable](s []T) []T {
	return lo.Compact(s)
}

// RemoveZeroValues remove zero values from a generic slice
/*
PS: 使用了反射，性能可能有问题，要避免大量调用!!!

@return 可能为nil
*/
func RemoveZeroValues[T any](slice []T) []T {
	var result []T
	var zeroValue T

	for _, v := range slice {
		if !reflect.DeepEqual(v, zeroValue) {
			result = append(result, v)
		}
	}
	return result
}
