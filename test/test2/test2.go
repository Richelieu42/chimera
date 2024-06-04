package main

import (
	"errors"
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"net"
)

func main() {
	var err error = &net.ParseError{
		Type: "Type",
		Text: "Text",
	}

	var c net.Error
	if errors.As(err, &c) {
		fmt.Println(c.Timeout())
	}
	fmt.Println("===")
}
