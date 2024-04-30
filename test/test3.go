package main

import (
	"github.com/centrifugal/gocent/v3"
)

func main() {
	client := gocent.New(gocent.Config{
		Addr:       "",
		GetAddr:    nil,
		Key:        "",
		HTTPClient: nil,
	})
}
