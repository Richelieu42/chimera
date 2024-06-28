package main

import (
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"os"
)

func main() {
	zapKit.NewWriteSyncerWithLock(os.Stdout)
}
