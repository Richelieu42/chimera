package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	log.Println("hello world")

	sb := &strings.Builder{}
	sb.WriteString("hello\n")
	sb.WriteString("world\n")
	fmt.Println(sb.String())
}
