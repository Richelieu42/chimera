package main

import (
	"fmt"
	"strings"
)

func main() {
	sb := &strings.Builder{}

	sb.WriteString("hello\n")
	sb.WriteString("world\n")
	fmt.Println(sb.String())
}
