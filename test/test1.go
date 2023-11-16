package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/list"
)

func main() {
	l := list.NewWriter()
	l.AppendItem("Game Of Thrones")
	l.Indent()
	l.AppendItems([]interface{}{"Winter", "Is", "Coming"})
	l.Indent()
	l.AppendItems([]interface{}{"This", "Is", "Known"})
	l.UnIndent()
	l.UnIndent()
	l.AppendItem("The Dark Tower")
	l.Indent()
	l.AppendItem("The Gunslinger")

	fmt.Println(l.Render())
	fmt.Println("------")

	l.SetStyle(list.StyleConnectedRounded)
	fmt.Println(l.Render())
}
