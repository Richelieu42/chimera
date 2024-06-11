package reqKit

import (
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

		 */
		Timeout time.Duration

		// InsecureSkipVerify
		/*
			true:  跳过证书验证
			false: 不跳过证书验证（默认; 更加安全）
		*/
		InsecureSkipVerify bool
	}

	ClientOption func(*clientOptions)
)

func loadClientOptions(options ...ClientOption) *clientOptions {
	opts := &clientOptions{
		Timeout:            DefaultTimeout,
		InsecureSkipVerify: false,
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
