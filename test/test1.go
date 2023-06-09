package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

func main() {
	fmt.Println(funcKit.AddEntireCaller(1, "123"))

	fmt.Printf("%+v\n", a())
	fmt.Println("=========")
	fmt.Printf("%+v\n", b())
	fmt.Println("=========")
	fmt.Printf("%+v\n", c())
}

func a() error {
	return errorKit.New("123")
}

func b() error {
	str := "   "
	if err := strKit.AssertNotBlank(str, "str"); err != nil {
		return err
	}
	return nil
}

func c() error {
	return errorKit.Wrap(redis.Nil, "123")
}
