package main

import (
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := jsonKit.AssertJsonString("[{}", "jsonStr"); err != nil {
		logrus.Error(err.Error())
	}
}
