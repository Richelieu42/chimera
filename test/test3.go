package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/compress/gzipKit"
)

func main() {
	json := `{"method":1,"data":{"text":"[127.0.0.1:8888, unknown, pulsar] Hello, id of this channel is [coc8tnh97i60icb4ajs0]."}}`
	fmt.Println(len(json))

	{
		data, err := gzipKit.Gzip([]byte(json), 1)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(data))
	}

	{
		data, err := gzipKit.Gzip([]byte(json), 9)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(data))
	}
}
