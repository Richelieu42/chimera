package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/component/web/request/reqKit"
	"time"
)

func main() {
	urls := []string{
		"http://127.0.0.1:8000/test",
		"http://127.0.0.1:8001/test",
		"http://127.0.0.1:8002/test",
		"http://127.0.0.1:8003/test",
	}

	c := reqKit.NewClient(reqKit.WithDev())
	lbc, err := reqKit.NewLbClient(c, urls, time.Millisecond*100, nil)
	if err != nil {
		panic(err)
	}
	resp, err := lbc.Get(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ToString())
}
