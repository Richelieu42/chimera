package sliceKit

import (
	"fmt"
	"testing"
)

func TestRemoveByIndex(t *testing.T) {
	s := []int{0, 1, 2}
	s1, item, ok := RemoveByIndex(s, 2)

	fmt.Println(s)    // [0 1 2]
	fmt.Println(s1)   // [0 1]
	fmt.Println(item) // 2
	fmt.Println(ok)   // true
}

func TestRemove(t *testing.T) {
	s := []int{0, 1, 0}
	fmt.Println(Remove(s, 0)) // [1 0] true
}
