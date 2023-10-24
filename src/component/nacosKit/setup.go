package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/richelieu-yang/chimera/v2/src/copyKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/intKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
	"github.com/sirupsen/logrus"
	"net/url"
)

var (
	clientConfig  *constant.ClientConfig
	serverConfigs []constant.ServerConfig
)

func MustSetUp(config *Config, options ...constant.ClientOption) {
	err := SetUp(config, options...)
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// SetUp
/*
@param options		!!!:
					(1) 建议配置 客户端的缓存目录（default is current path）		constant.WithCacheDir
					(2) 建议配置 客户端的日志目录（default is current path）		constant.WithLogDir
					(3) 建议配置 客户端的日志级别（default value is info）		constant.WithLogLevel	"debug" || "info"（默认）...
*/
func SetUp(config *Config, options ...constant.ClientOption) (err error) {
	defer func() {
		if err != nil {
			clientConfig = nil
			serverConfigs = nil
		}
	}()

	/* (0) validate */
	if err = validateKit.Struct(config); err != nil {
		err = errorKit.Wrap(err, "Fail to verify")
		return
	}

	/* (1) clientConfig */
	options1 := []constant.ClientOption{
		constant.WithNamespaceId(config.NamespaceId),
		constant.WithBeatInterval(5000), // 5000ms
		constant.WithNotLoadCacheAtStart(true),
	}
	options1 = append(options1, options...)
	clientConfig = constant.NewClientConfig(options1...)

	/* (2) serverConfigs */
	if err = sliceKit.AssertNotEmpty(config.Addrs, "config.Addrs"); err != nil {
		return
	}
	for _, addr := range config.Addrs {
		if strKit.IsBlank(addr) {
			continue
		}

		var u *url.URL
		u, err = urlKit.Parse(addr)
		if err != nil {
			err = errorKit.Wrap(err, "Fail to parse address(%s).", addr)
			return
		}
		var port uint64
		port, err = intKit.ToUint64E(u.Port())
		if err != nil {
			err = errorKit.Wrap(err, "Invalid address(%s) with port string(%s).", addr, u.Port())
			return
		}
		if err = netKit.AssertValidPort(int(port)); err != nil {
			err = errorKit.Wrap(err, "Invalid address(%s) with port(%d).", addr, port)
			return
		}

		serverConfigs = append(serverConfigs, constant.ServerConfig{
			Scheme:      u.Scheme,
			IpAddr:      u.Hostname(),
			Port:        port,
			ContextPath: u.Path,
		})
	}
	if sliceKit.IsEmpty(serverConfigs) {
		err = errorKit.New("No valid address.")
		return
	}

	/* (3) test（以防配置有问题） */
	cc, err := NewConfigClient()
	if err != nil {
		return
	}
	defer cc.CloseClient()
	nc, err := NewNamingClient()
	if err != nil {
		return
	}
	defer nc.CloseClient()

	return
}

func GetClientConfigCopy(options ...constant.ClientOption) (*constant.ClientConfig, error) {
	if clientConfig == nil || serverConfigs == nil {
		return nil, NotSetUpError
	}

	// 深拷贝
	clientConfig1, err := copyKit.DeepCopy(clientConfig)
	if err != nil {
		return nil, errorKit.Wrap(err, "Fail to deep copy.")
	}

	// 再次修改 *constant.ClientConfig
	for _, option := range options {
		option(clientConfig1)
	}

	return clientConfig, nil
}

// GetNamespaceId
/*
PS: 如果一个服务仅使用一个NamespaceId，可以调用此方法.
*/
func GetNamespaceId() (string, error) {
	if clientConfig == nil || serverConfigs == nil {
		return "", NotSetUpError
	}

	return clientConfig.NamespaceId, nil
}
