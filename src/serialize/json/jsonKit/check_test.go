package jsonKit

import (
	"fmt"
	"testing"
)

func TestIsJsonString(t *testing.T) {
	fmt.Println(IsJsonString(""))
	fmt.Println(IsJsonString("{}"))
	fmt.Println(IsJsonString("[]"))
	fmt.Println(IsJsonString("qdqwd"))
	fmt.Println(IsJsonString("{]"))
}
