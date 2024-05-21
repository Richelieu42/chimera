package main

import "github.com/richelieu-yang/chimera/v3/src/core/mapKit"

type bean struct {
	A int `json:"a"`
}

func main() {
	m := map[string]interface{}{
		"a": 0,
	}

	var b *bean
	mapKit.Decode(m, b)

}
