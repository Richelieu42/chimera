package main

import (
	"fmt"
	"github.com/imroc/req/v3"
	"net/http"
)

func main() {
	c := req.C().
		DevMode().
		SetCommonPathParam("a", "test").
		SetBaseURL("http://127.0.0.1:8001").
		AddCommonQueryParams("a", "0", "1").
		SetCommonHeaderNonCanonical("Yozo-Source", "1").
		SetCommonCookies(&http.Cookie{Name: "f", Value: "v"}, &http.Cookie{Name: "f1", Value: "v1"})

	// 此时真正的请求url为 http://127.0.0.1:8001/test?a=0&a=1
	resp := c.Post("/{a}").
		SetFile("file", "_upload.txt").
		Do()
	if resp.Err != nil {
		panic(resp.Err)
	}

	fmt.Println(resp.Request.URL.String()) // http://127.0.0.1:8001/test?a=0&a=1
	fmt.Println(resp.Status)               // 200 OK
	fmt.Println(resp.ToString())           // Hello world! <nil>
}
