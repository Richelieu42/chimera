package dataSizeKit

import (
	"fmt"
	"testing"
)

func TestByteToMiB(t *testing.T) {
	bytes := MiB*10 + KiB*123

	fmt.Println(ByteToMiB(bytes, 2)) // 10.12
	fmt.Println(ByteToMiB(bytes, 3)) // 10.12
	fmt.Println(ByteToMiB(bytes, 4)) // 10.1201
}
