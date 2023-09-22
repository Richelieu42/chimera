package skyWalkingKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
	"github.com/sirupsen/logrus"
)

// MustSetUp
/*
TODO: 待实现.
*/
func MustSetUp(config *Config) {
	if err := SetUp(config); err != nil {
		logrus.Fatal(err)
	}
}

// SetUp
/*
TODO: 待实现.
*/
func SetUp(config *Config) error {
	serverAddr := config.ServerAddr
	addr, err := netKit.ParseToAddress(serverAddr)
	if err != nil {
		return err
	}
	fmt.Println(addr.String())

	return nil
}
