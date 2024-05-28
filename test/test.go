package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
)

func main() {
	//if err := jsonKit.AssertJsonString("[{}", "jsonStr"); err != nil {
	//	logrus.Error(err.Error())
	//}

	err := errorKit.Simplef("111")
	fmt.Println(err.Error())

	fmt.Println("===")

	fmt.Printf("%v\n", err)

	fmt.Println("===")

	fmt.Printf("%+v\n", err)
}
