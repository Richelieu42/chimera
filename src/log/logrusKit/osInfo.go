package logrusKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/osKit"
	"github.com/sirupsen/logrus"
)

func printOsInfo() {
	logrus.Infof("[CHIMERA, OS] os: [%s].", osKit.OS)
	logrus.Infof("[CHIMERA, OS] arch: [%s].", osKit.ARCH)
	logrus.Infof("[CHIMERA, OS] bits: [%d].", osKit.GetOsBits())

	printUniqueOsInfo()
}
