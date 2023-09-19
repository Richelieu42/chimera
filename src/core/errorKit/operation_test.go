package errorKit

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestIs(t *testing.T) {
	err := redis.Nil
	err1 := Wrap(err, "1")
	err2 := Wrap(err1, "2")

	fmt.Printf("%+v\n", err2)

	fmt.Println(Is(err2, err)) // true
	fmt.Println(Is(err1, err)) // true
	fmt.Println(Is(err, err))  // true

	fmt.Println(Is(err2, err1)) // true
	fmt.Println(Is(err1, err2)) // false
}

type myError struct {
	Text string
}

func (err myError) Error() string {
	return err.Text
}

// TestAs receiver为"值类型"
func TestAs(t *testing.T) {
	err := myError{
		Text: "cyy",
	}
	err1 := Wrap(err, "1")

	target := &myError{}
	if ok := As(err1, target); !ok {
		panic("ok == false")
	}
	fmt.Println(target.Text) // cyy
	if err.Text != target.Text {
		panic("not equal")
	}
}

type myError1 struct {
	Text string
}

func (err *myError1) Error() string {
	return err.Text
}

// TestAs1 receiver为"指针类型"
func TestAs1(t *testing.T) {
	err := &myError1{
		Text: "cyy",
	}
	err1 := Wrap(err, "1")

	target := &myError1{}
	// 相较于 TestAs，多了个 &
	if ok := As(err1, &target); !ok {
		panic("ok == false")
	}
	fmt.Println(target.Text) // cyy
	if err.Text != target.Text {
		panic("not equal")
	}
}
