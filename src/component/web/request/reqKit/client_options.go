package reqKit

import (
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/core/mathKit"
	"time"
)

const (
	// MinTimeout 最小超时时间
	MinTimeout = time.Second * 3

	// DefaultTimeout 默认超时时间
	DefaultTimeout = time.Second * 30
)

type (
	clientOptions struct {
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

func loadClientOptions(baseClient *req.Client, options ...ClientOption) *clientOptions {
	opts := &clientOptions{
		Timeout:            DefaultTimeout,
		InsecureSkipVerify: false,
		// imroc/req默认: 输出到 os.Stdout
		Logger: baseClient.GetLogger(),
	}

	for _, option := range options {
		option(opts)
	}

	opts.Timeout = mathKit.Max(opts.Timeout, MinTimeout)

	return opts
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

func WithLogger(logger req.Logger) ClientOption {
	return func(options *clientOptions) {
		options.Logger = logger
	}
}
