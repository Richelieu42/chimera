package jsonFieldKit

import (
	"fmt"
	"testing"
)

func TestGetInt64SliceFieldFromString(t *testing.T) {
	jsonStr := `{"b":{"c":[1,2,3,4,5]}}`
	fmt.Println(GetInt64SliceFieldFromString(jsonStr, "b.c")) // [1 2 3 4 5]
}
