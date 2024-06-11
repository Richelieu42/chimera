package main

import (
	"fmt"
	"github.com/imroc/req/v3"
)

func main() {
	client := req.C()

	tp := client.GetTransport()
	fmt.Println(tp)
	fmt.Println(tp.TLSClientConfig)
	fmt.Println("InsecureSkipVerify:", tp.TLSClientConfig.InsecureSkipVerify)
}
