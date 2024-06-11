package main

import (
	"fmt"
	"github.com/imroc/req/v3"
)

type APIResponse struct {
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

func main() {
	var resp APIResponse
	c := req.C()

	c.R().Do()

	err := c.Post("https://httpbin.org/post"). // method + url
							SetBody("hello"). // set request body
							Do().             // send request
							Into(&resp)       // unmarshal response body
	if err != nil {
		panic(err)
	}
	fmt.Println("My IP is", resp.Origin)
}
