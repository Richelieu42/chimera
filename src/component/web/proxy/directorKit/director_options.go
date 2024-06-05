package directorKit

type (
	directorOptions struct {
		// scheme "http"（默认） || "https"
		scheme string

		// requestUrlPath 请求路由
		requestUrlPath *string

		// overrideQueryParams （优先级高于extraQueryParams）不会保留原先的queries.
		/*
			PS: 切片中的字符串应当是 未处理（编码） 过的.
		*/
		overrideQueryParams map[string][]string

		// extraQueryParams 会保留原先的queries.
		/*
			PS: 切片中的字符串应当是 未处理（编码） 过的.
		*/
		extraQueryParams map[string][]string
	}

	DirectorOption func(opts *directorOptions)
)

func loadOptions(options ...DirectorOption) *directorOptions {
	opts := &directorOptions{
		scheme:              "",
		requestUrlPath:      nil,
		overrideQueryParams: nil,
		extraQueryParams:    nil,
	}

	for _, option := range options {
		option(opts)
	}
	// 默认值
	if opts.scheme == "" {
		opts.scheme = "http"
	}
	return opts
}

func WithScheme(scheme string) DirectorOption {
	return func(opts *directorOptions) {
		opts.scheme = scheme
	}
}

func WithRequestUrlPath(requestUrlPath string) DirectorOption {
	return func(opts *directorOptions) {
		opts.requestUrlPath = &requestUrlPath
	}
}

func WithOverrideQueryParams(overrideQueryParams map[string][]string) DirectorOption {
	return func(opts *directorOptions) {
		opts.overrideQueryParams = overrideQueryParams
	}
}

func WithExtraQueryParams(extraQueryParams map[string][]string) DirectorOption {
	return func(opts *directorOptions) {
		opts.extraQueryParams = extraQueryParams
	}
}
