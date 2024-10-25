package main

import (
	"github.com/richelieu-yang/chimera/v3/src/log/console"
)

func main() {
	console.Info("info")
	console.Info("warn")
	console.Fatal("fatal")
	console.Error("error")
}
