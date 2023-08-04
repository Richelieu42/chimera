package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
)

// NewClient
/*
@param logPath 客户端的日志输出（为空则输出到控制台; 不会rotate）
*/
func NewClient(addresses []string, logPath string) (pulsar.Client, error) {
	/* url */
	tmp, err := netKit.ProcessAddresses(addresses)
	if err != nil {
		return nil, err
	}
	url := UrlPrefix + sliceKit.Join(tmp, ",")

	/* logger */
	var logger log.Logger
	if strKit.IsNotEmpty(logPath) {
		fileLogger, err := logrusKit.NewFileLogger(logPath)
		if err != nil {
			return nil, err
		}
		logger = log.NewLoggerWithLogrus(fileLogger)
	}

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:    url,
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
