package qrCodeKit

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	content, err := Read("read_demo.png")
	if err != nil {
		panic(err)
	}
	fmt.Println("content:", content) // content: Hello World
}
