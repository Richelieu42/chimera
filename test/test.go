package main

import (
	"fmt"
)

func main() {
	var m map[string]interface{}
	fmt.Println(len(m))

	var s []string = nil
	var s1 []string = []string{}

	fmt.Println(len(s))
	fmt.Println(len(s1))
}
