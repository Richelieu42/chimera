package qrCodeKit

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	content, err := Decode("_test2.png")
	if err != nil {
		panic(err)
	}
	fmt.Println("content:", content)
}
