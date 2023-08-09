package logrusKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/osKit"
	"github.com/sirupsen/logrus"
)

// printUniqueOsInfo 输出特殊的信息（主要依赖于不同的OS）
func printUniqueOsInfo() {
	if str, err := osKit.GetUlimitInfo(); err != nil {
		logrus.WithError(err).Error("[CHIMERA, OS] fail to get ulimit information")
	} else {
		DisableQuoteTemporarily(nil, func(logger *logrus.Logger) {
			logger.Infof("[CHIMERA, OS] ulimit information:\n%s\n", str)
		})
	}
}
