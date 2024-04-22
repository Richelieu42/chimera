package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/crypto/base64Kit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
)

func main() {
	fmt.Println(base64Kit.DecodeStringToString("dGVzdCDmtYvor5U="))
}
