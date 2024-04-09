package main

import (
	"embed"
	"fmt"
	"io"
)

//go:embed c
var a embed.FS

//go:embed c/a.json
var b []byte

//go:embed c/a.json
var c string

func main() {
	f, err := a.Open("c/a.json")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	fmt.Println(string(b))

	fmt.Println(c)
}
