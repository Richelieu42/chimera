package main

import (
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
)

func main() {
	json, err := jsonKit.MarshalToString(1)
	if err != nil {
		panic(err)
	}
	console.Info(json)
}
