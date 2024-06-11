package main

import (
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

type APIResponse struct {
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

func main() {
	var logger req.Logger = logrus.StandardLogger()
	fmt.Println(logger)

	var resp APIResponse
	c := req.C()

	c.SetScheme()
	c.SetProxyURL()
	c.SetBaseURL()

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
