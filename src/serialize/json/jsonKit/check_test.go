package jsonKit

import (
	"fmt"
	"testing"
)

func TestIsJsonString(t *testing.T) {
	fmt.Println(IsJsonString(""))      // false
	fmt.Println(IsJsonString("{}"))    // true
	fmt.Println(IsJsonString("[]"))    // true
	fmt.Println(IsJsonString("qdqwd")) // false
	fmt.Println(IsJsonString("{]"))    // false
}

func TestIsJson(t *testing.T) {
	fmt.Println(IsJson([]byte("")))      // false
	fmt.Println(IsJson([]byte("{}")))    // true
	fmt.Println(IsJson([]byte("[]")))    // true
	fmt.Println(IsJson([]byte("qdqwd"))) // false
	fmt.Println(IsJson([]byte("{]")))    // false
}
