package sliceKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/runtimeKit"
	"testing"
)

func TestClear(t *testing.T) {
	type bean struct {
		id int
	}

	fmt.Println(runtimeKit.GoVersion)

	s := make([]bean, 3, 6)
	s[0] = bean{0}
	s[1] = bean{1}
	s[2] = bean{2}
	fmt.Println(s, len(s), cap(s)) // [{0} {1} {2}] 3 6

	Clear(s)
	fmt.Println(s, len(s), cap(s)) // [{0} {0} {0}] 3 6

	s[0].id = 999
	fmt.Println(s, len(s), cap(s)) // [{999} {0} {0}] 3 6
}
