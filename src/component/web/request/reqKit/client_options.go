package reqKit

import "time"

const (
	MinTimeout = time.Second * 3
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
		Timeout:            0,
		InsecureSkipVerify: false,
	}

	for _, option := range options {
		option(opts)
	}

	if opts.Timeout < MinTimeout {
		opts.Timeout = MinTimeout
	}

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
