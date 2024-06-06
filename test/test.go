package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
)

func main() {
	str := `\"{\\\"jsonParams\\\":{\\\"method\\\":95,\\\"fileId\\\":\\\"ff3f56f238364821ab4beb5b68f4d45945243258\\\",\\\"params\\\":{\\\"operId\\\":\\\"192.168.134.1$0\\\",\\\"fileId\\\":\\\"ff3f56f238364821ab4beb5b68f4d45945243258\\\"}},\\\"fileId\\\":\\\"ff3f56f238364821ab4beb5b68f4d45945243258\\\",\\\"header\\\":{\\\"Yozo-Authorization\\\":\\\"Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc3MjI0OTksImZmM2Y1NmYyMzgzNjQ4MjFhYjRiZWI1YjY4ZjRkNDU5NDUyNDMyNTgiOiIxOTIuMTY4LjEzNC4xIn0.S3UFYl6lN2uKEzOOI9WCy1-JQhfInXrYqQ3jLzQdXRTpvExipXshY5cNX_d1lSNc5pZogUzvHbxOEmbaMgRQVw\\\"},\\\"target\\\":\\\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJNYW5hZ2VyIiwiaG9zdCI6IjE5Mi4xNjguMTM0LjEiLCJodHRwUG9ydCI6MTAwMDAsImdycGNQb3J0IjoxOTAwMH0.XnvArmQ2uha-ib48OAFbnmrY5IuVz19dXGBlConPe2o\\\"}\"`
	var m map[string]interface{}
	if err := jsonKit.UnmarshalFromString(str, &m); err != nil {
		panic(err)
	}
	fmt.Println(jsonKit.MarshalIndentToString(m, "", "    "))
}
