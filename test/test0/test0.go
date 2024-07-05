package main

import (
	"github.com/imroc/req/v3"
)

func main() {
	client := req.C().EnableInsecureSkipVerify()

	client

}
