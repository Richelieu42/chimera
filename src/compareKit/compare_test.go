package compareKit

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	type bean struct {
		Name   string
		Lovers []string
	}
	b0 := &bean{
		Name:   "张三",
		Lovers: []string{"李四"},
	}
	b1 := &bean{
		Name:   "张三",
		Lovers: []string{"李四"},
	}
	fmt.Println(Equal(b0, b1))
}

type bean struct {
	Name   string
	Lovers []string
}

func (b *bean) Equal(b1 *bean) bool {
	return b.Name == b1.Name
}

// 结构体实现了: (T) Equal(T) bool 或者 (T) Equal(I) bool
func TestEqual1(t *testing.T) {
	b0 := &bean{
		Name:   "张三",
		Lovers: []string{"李四"},
	}
	b1 := &bean{
		Name:   "张三",
		Lovers: []string{"李4"},
	}
	fmt.Println(Equal(b0, b1))
}
