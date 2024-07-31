package sliceKit

import "github.com/samber/lo"

// IndexOf
/*
@param s 可以为nil（此时返回-1）
@return -1: 传参ele不存在于传参s 中
*/
func IndexOf[T comparable](s []T, ele T) int {
	return lo.IndexOf(s, ele)
}

// Find
/*
@param s			可以为nil
@param predicate	返回true: 该元素是要找的
@return 			第2个返回值: 是否找到？
*/
func Find[T any](s []T, predicate func(item T) bool) (T, bool) {
	return lo.Find(s, predicate)
}

// FindOrElse
/*
@param s			可以为nil
@param fallback		默认值（找不到就返回它）
@param predicate	返回true: 该元素是要找的
*/
func FindOrElse[T any](s []T, fallback T, predicate func(item T) bool) T {
	return lo.FindOrElse(s, fallback, predicate)
}

// LastIndexOf
/*
@param s 	可以为nil（此时返回-1）
@return 	如果不存在于切片实例中的话，返回-1
*/
func LastIndexOf[T comparable](s []T, element T) int {
	return lo.LastIndexOf(s, element)
}

// FindIndexOf
/*
@param s			可以为nil
@param predicate	返回true: 该元素是要找的
@param 				如果找不到，第2个返回值为-1
*/
func FindIndexOf[T any](s []T, predicate func(item T) bool) (T, int, bool) {
	return lo.FindIndexOf(s, predicate)
}

// FindLastIndexOf
/*
@param s			可以为nil
@param predicate	返回true: 该元素是要找的
@param 				如果找不到，第2个返回值为-1
*/
func FindLastIndexOf[T any](s []T, predicate func(item T) bool) (T, int, bool) {
	return lo.FindLastIndexOf(s, predicate)
}
