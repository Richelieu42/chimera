package main

import (
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(logrus.StandardLogger().Level.String())
}
