package ginKit

type (
	Config struct {
		Mode     string `json:"mode" yaml:"mode" validate:"oneof=debug release test"`
		HostName string `json:"hostName" yaml:"hostName" validate:"omitempty,hostname|ipv4"`
		// Port
		/*
			0（默认）: 不使用 http port
		*/
		Port     int  `json:"port" yaml:"port" validate:"port|eq=0,necsfield=SSL.Port"`
		Colorful bool `json:"colorful" yaml:"colorful"`
		Pprof    bool `json:"pprof" yaml:"pprof"`

		SSL        SslConfig        `json:"ssl" yaml:"ssl"`
		Middleware MiddlewareConfig `json:"middleware" yaml:"middleware"`
	}

	SslConfig struct {
		// Port
		/*
			0（默认）: 不使用 https port
		*/
		Port     int    `json:"port" yaml:"port" validate:"port|eq=0"`
		CertFile string `json:"certFile" yaml:"certFile" validate:"file_unless=Port 0"`
		KeyFile  string `json:"keyFile" yaml:"keyFile" validate:"file_unless=Port 0"`
	}

	MiddlewareConfig struct {
		BodyLimit     int64  `json:"bodyLimit" yaml:"bodyLimit"`
		Gzip          bool   `json:"gzip" yaml:"gzip"`
		XFrameOptions string `json:"xFrameOptions" yaml:"xFrameOptions" validate:"omitempty,lowercase,oneof=deny sameorigin|startswith=allow-from "`

		Cors CorsConfig `json:"cors" yaml:"cors"`
		//Referer       []*refererKit.RefererVerifierBuilder `json:"referer" yaml:"referer"`
		RateLimiter *RateLimiterConfig
	}

	// CorsConfig cors（跨源资源共享）的配置
	CorsConfig struct {
		Access  bool     `json:"access" yaml:"access"`
		Origins []string `json:"origins" yaml:"origins" validate:"unique,dive,required"`
	}

	// RateLimiterConfig 限流器的配置（令牌桶算法）
	RateLimiterConfig struct {
		Limit int `json:"limit" yaml:"limit" validate:"gt=0"`
		Burst int `json:"burst" yaml:"burst" validate:"gt=0,gtecsfield=Limit"`
	}
)
