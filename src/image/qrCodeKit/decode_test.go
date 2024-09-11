package qrCodeKit

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	content, err := Decode("_read_demo.png")
	if err != nil {
		panic(err)
	}
	fmt.Println("content:", content) // content: Hello World
}
