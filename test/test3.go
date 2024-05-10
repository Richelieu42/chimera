package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	defer func() {
		if obj := recover(); obj != nil {
			err, ok := obj.(error)
			fmt.Println("ok", ok)
			if ok {
				fmt.Println(errors.Is(err, http.ErrAbortHandler))
			}
		}
	}()

	test()
}

func test() {
	panic(http.ErrAbortHandler)
}
