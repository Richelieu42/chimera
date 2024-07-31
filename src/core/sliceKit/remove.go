package sliceKit

// RemoveByIndex 根据下标删除指定元素.
/*
PS:
(0) 也可以用于移除"第一个"或"最后一个"个元素；
(1) 为什么要有两个返回值且第一个是slice？因为 append() 并不会修改传参slice，返回的是一个新的slice.
(2) 调用此方法后，必须第一时间将 第一个返回值 赋值给 传进来的那个slice变量！！！

@param s		(1) 可以为nil
				(2) 不会修改传参s
@param index 	下标（索引），取值范围: [0, length)

e.g.
	s := []int{0, 1, 2}
	s1, item, ok := RemoveByIndex(s, 2)

	fmt.Println(s)    // [0 1 2]
	fmt.Println(s1)   // [0 1]
	fmt.Println(item) // 2
	fmt.Println(ok)   // true
*/
func RemoveByIndex[T any](s []T, index int) (rst []T, item T, ok bool) {
	if len(s) == 0 {
		rst = s
		return
	}

	item = s[index]

	// !!!: 下面一行代码执行后，会修改外部的slice
	//rst = append(s[:index], s[index+1:]...)

	rst = append(rst, s[:index]...)
	rst = append(rst, s[index+1:]...)

	ok = true
	return
}

// RemoveFirst 移除第一个元素（如果有的话）
/*
@param s 可以为nil

e.g.
[int](nil) 		=> [] 0 false
([]int{0}) 		=> [] 0 true
([]int{0, 1}) 	=> [1] 0 true
*/
func RemoveFirst[T any](s []T) (s1 []T, item T, ok bool) {
	if len(s) == 0 {
		s1 = s
		return
	}

	// 此时第一个返回值 必定不为nil && len >= 0
	item = s[0]
	s1 = s[1:]
	ok = true
	return
}

// RemoveLast 移除最后一个元素（如果有的话）
/*
@param s 可以为nil

e.g.
[int](nil) 		=> [] 0 false
([]int{0}) 		=> [] 0 true
([]int{0, 1}) 	=> [0] 1 true
*/
func RemoveLast[T any](s []T) (s1 []T, item T, ok bool) {
	if len(s) == 0 {
		s1 = s
		return
	}

	// 此时第一个返回值 必定不为nil && len >= 0
	index := len(s) - 1
	item = s[index]
	s1 = s[0:index]
	ok = true
	return
}

// Remove 移除元素.
/*
PS:
(1) 切片实例s中，存在多个item的话，仅会移除第一个.

@param s (1)可以为nil (2)不会修改传参s

e.g.	反例
	texts := []string{"0", "1", "2"}
	fmt.Println(texts) // [0 1 2]

	texts1, _ := sliceKit.Remove(texts, "1")
	fmt.Println(texts)  // [0 2 2]
	fmt.Println(texts1) // [0 2]
*/
func Remove[T comparable](s []T, item T) ([]T, bool) {
	index := IndexOf(s, item)
	if index == -1 {
		// 如果元素不存在于切片中，返回原始切片
		return s, false
	}

	s, _, ok := RemoveByIndex(s, index)
	return s, ok
}

// RemoveBy 移除不满足条件的元素（返回的是一个新的slice实例）.
/*
@param s 			可以为nil
@param predicate	(1)不能为nil (2)返回值为true: 移除当前元素
@return (1)非nil (2)len>=0
*/
func RemoveBy[T comparable](s []T, predicate func(element T) bool) []T {
	result := make([]T, 0, len(s))

	for _, element := range s {
		if !predicate(element) {
			result = append(result, element)
		}
	}
	return result
}
