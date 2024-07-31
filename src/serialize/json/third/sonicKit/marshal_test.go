package sonicKit

import (
	"fmt"
	"testing"
)

func TestMarshalToStringByAPIWithIndent(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": 2,
		"":  3,
		"d": 4,
	}
	str, err := MarshalToStringByAPIWithIndent(nil, m, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
