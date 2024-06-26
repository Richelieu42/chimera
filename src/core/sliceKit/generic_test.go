package sliceKit

import (
	"fmt"
	"testing"
)

func TestRemoveZeroValues(t *testing.T) {
	s0 := []int{0, 1, 2, 0, 3, 0, 4}
	s1 := []string{"", "hello", "", "world", ""}
	s2 := []interface{}{nil, nil, 1, false, 0, "", true}

	fmt.Println(RemoveZeroValues(s0)) // [1 2 3 4]
	fmt.Println(RemoveZeroValues(s1)) // [hello world]
	fmt.Println(RemoveZeroValues(s2)) // [1 false 0  true]
}
