package snappyKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/bytesKit"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	src := []byte("11111111222222222223333333333444444444444444444444444")
	//src := []byte("akakakakakakakakakakdddddddddcccccceeeeeeeegggggggggsssss")

	dst := Encode(src)
	fmt.Println(string(dst), len(dst))

	src1, err := Decode(dst)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(src1), len(src1))

	if !bytesKit.Equals(src, src1) {
		panic("not equal")
	}
	fmt.Println("equal")
}
