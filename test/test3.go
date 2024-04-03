package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
)

func main() {
	type config struct {
		Level *int
	}

	c := &config{}
	_, err := viperKit.UnmarshalFromFile("test.yaml", nil, c)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
