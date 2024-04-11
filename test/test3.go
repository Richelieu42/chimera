package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"

	"github.com/rifflock/lfshook"
)

func main() {
	hook := lfshook.NewHook()
}
