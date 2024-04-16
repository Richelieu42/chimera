package gzipKit

import (
	"fmt"
	"testing"
)

func TestCompressAndDecompress(t *testing.T) {
	json := `===
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
Hello, World! This is a sample byte slice to be compressed using LZ4.
---`
	fmt.Println("before length:", len(json))
	fmt.Println("------")

	{
		// 压缩
		data, err := Compress([]byte(json), WithCompressThreshold(-1))
		if err != nil {
			panic(err)
		}
		fmt.Println("after length:", len(data))
		fmt.Println(string(data))
		fmt.Println("------")
	}

	{
		// 不压缩（497 < 500）
		data, err := Compress([]byte(json), WithCompressThreshold(500))
		if err != nil {
			panic(err)
		}
		fmt.Println("after length:", len(data))
		fmt.Println(string(data))
		fmt.Println("------")
	}
}
