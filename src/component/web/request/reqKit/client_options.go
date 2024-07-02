package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"time"
)

const (
	// DefaultTimeout 默认超时时间
	DefaultTimeout = time.Second * 30
)

type (
	clientOptions struct {
		Dev bool

		// Timeout
		/*
			超时时间（发送请求的整个周期，includes connection time, any redirects, and reading the response body）
		*/
		Timeout time.Duration

		// InsecureSkipVerify
		/*
			true:  跳过证书验证
			false: 不跳过证书验证（默认; 更加安全）
		*/
		InsecureSkipVerify bool

		// Logger 日志输出
		/*
			nil: 不输出
		*/
		Logger req.Logger
	}

	ClientOption func(*clientOptions)
)

func loadClientOptions(options ...ClientOption) *clientOptions {
	logger := zapKit.NewLogger(nil).Sugar()

	opts := &clientOptions{
		Dev:                false,
		Timeout:            DefaultTimeout,
		InsecureSkipVerify: true,
		// imroc/req默认: 输出到 os.Stdout
		Logger: logger,
	}

	for _, option := range options {
		option(opts)
	}

	if opts.Timeout <= 0 {
		opts.Timeout = DefaultTimeout
	}

	return opts
}

func WithDev() ClientOption {
	return func(options *clientOptions) {
		options.Dev = true
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(options *clientOptions) {
		options.Timeout = timeout
	}
}

func WithInsecureSkipVerify(insecureSkipVerify bool) ClientOption {
	return func(options *clientOptions) {
		options.InsecureSkipVerify = insecureSkipVerify
	}
}

// WithLogger
/*
@param logger	(1) 可以是 *logrus.Logger 实例 || *zap.SugaredLogger 实例
				(2) 可以为nil（disable log, 禁用输出）
*/
func WithLogger(logger req.Logger) ClientOption {
	return func(options *clientOptions) {
		options.Logger = logger
	}
}
