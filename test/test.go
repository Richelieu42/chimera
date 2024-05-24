package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
)

func main() {
	//hosts := []string{"127.0.0.1:80"}
	hosts := []string{"127.0.0.1:80", "127.0.0.1:81"}
	//err := validateKit.Var(hosts, "required,gte=2,unique,dive,hostname_port")
	err := validateKit.Var(hosts, "required,gte=1,unique,dive,hostname_port")
	if err != nil {
		panic(err)
	}
	fmt.Println("Pass")
}
