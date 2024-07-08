package main

import (
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"go.uber.org/zap"
)

func main() {
	console.Info("hello world", zap.String("c", ""), zap.String("a", ""), zap.String("b", ""))
}
