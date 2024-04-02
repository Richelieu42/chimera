package ioKit

import (
	"fmt"
	"testing"
)

func TestReadAtLeast(t *testing.T) {
	{
		reader := NewReaderFromString("abcdefghijklmn")
		buf := make([]byte, 3)

		fmt.Println(ReadAtLeast(reader, buf, 3)) // 3 <nil>
		fmt.Println(string(buf))                 // abc

		fmt.Println(ReadAtLeast(reader, buf, 3)) // 3 <nil>
		fmt.Println(string(buf))                 // def
	}

	{
		reader := NewReaderFromString("abcdefghijklmn")
		buf := make([]byte, 3)

		fmt.Println(ReadAtLeast(reader, buf, 4)) // 0 short buffer
	}
}
