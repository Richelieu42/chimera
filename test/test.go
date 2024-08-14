package main

import (
	"fmt"
	"github.com/samber/lo"
)

func main() {
	fmt.Println(lo.Substring("hello", -1, 1))
	//
	//result1 := lo.Substring("hello", 2, 3)
	//result2 := lo.Substring("hello", -4, 3)
	//result3 := lo.Substring("hello", -2, math.MaxUint)
	//
	//fmt.Printf("%v\n", result1)
	//fmt.Printf("%v\n", result2)
	//fmt.Printf("%v\n", result3)
}
