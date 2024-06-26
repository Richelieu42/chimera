package main

import "github.com/richelieu-yang/chimera/v3/src/log/zapKit"

func main() {
	zapKit.Info("hello")
	zapKit.Infof("hello %s", "world")
}
