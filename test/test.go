package main

import (
	"fmt"
	"reflect"
)

// RemoveZeroValues removes zero values from a generic slice
/*
PS: 使用了反射，性能可能有问题，要避免大量调用!!!
*/
func RemoveZeroValues[T comparable](slice []T) []T {
	var result []T
	var zeroValue T

	for _, v := range slice {
		if !reflect.DeepEqual(v, zeroValue) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	intSlice := []int{0, 1, 2, 0, 3, 0, 4}
	stringSlice := []string{"", "hello", "", "world", ""}

	fmt.Println(RemoveZeroValues(intSlice))    // Output: [1 2 3 4]
	fmt.Println(RemoveZeroValues(stringSlice)) // Output: [hello world]
}
