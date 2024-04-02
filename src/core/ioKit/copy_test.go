package ioKit

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	reader := NewReaderFromString("abcdefg")

	fmt.Println(Copy(buf, reader)) // 7 <nil>
	fmt.Println(buf.String())      // abcdefg

	/* 再次读取 */
	data, err := ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(data)) // 0（因为Copy操作会"读"reader）
}
