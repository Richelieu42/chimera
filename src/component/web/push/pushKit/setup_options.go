package pushKit

import (
	"time"
)

type (
	Options struct {
		WsPongInterval time.Duration

		SsePongInterval time.Duration
	}

	Option func(opts *Options)
)

func loadOptions(options ...Option) *Options {
	opts := &Options{
		WsPongInterval:  time.Second * 15,
		SsePongInterval: time.Second * 15,
	}

	for _, option := range options {
		option(opts)
	}
	return opts
}

func WithWsPongInterval(interval time.Duration) Option {
	return func(opts *Options) {
		opts.WsPongInterval = interval
	}
}

func WithSsePongInterval(interval time.Duration) Option {
	return func(opts *Options) {
		opts.SsePongInterval = interval
	}
}
