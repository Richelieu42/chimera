package sliceKit

import (
	"fmt"
	"testing"
)

func TestPolyfillStringSlice(t *testing.T) {
	s := []string{"", "1", " 1", "   "}
	PolyfillStringSlice(s)
	fmt.Println(s)
}
