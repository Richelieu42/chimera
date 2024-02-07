package zstdKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/bytesKit"
	"testing"
)

func TestCompressAndDecompress(t *testing.T) {
	data := []byte("hello world~!@#$%^&*()_+-=<>?,./;'[]{}|\\hello world~!@#$%^&*()_+-=<>?,./;'[]{}|\\hello world~!@#$%^&*()_+-=<>?,./;'[]{}|\\hello world~!@#$%^&*()_+-=<>?,./;'[]{}|\\hello world~!@#$%^&*()_+-=<>?,./;'[]{}|\\hello world~!@#$%^&*()_+-=<>?,./;'[]{}|\\hello world~!@#$%^&*()_+-=<>?,./;'[]{}|\\hello world~!@#$%^&*()_+-=<>?,./;'[]{}|\\hello world~!@#$%^&*()_+-=<>?,./;'[]{}|\\hello world~!@#$%^&*()_+-=<>?,./;'[]{}|\\")

	compressed, err := Compress(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(compressed))

	data1, err := Decompress(compressed)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data1))

	if !bytesKit.Equals(data, data1) {
		panic("not equals")
	}

	fmt.Println("len(data):", len(data))
	fmt.Println("len(compressed):", len(compressed))
}
