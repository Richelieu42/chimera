package ginKit

type (
	Config struct {
		Mode     string `json:"mode" yaml:"mode" validate:"oneof=debug release test"`
		HostName string `json:"hostName" yaml:"hostName"`
		// Port
		/*
			0（默认）: 不使用 http port
		*/
		Port     int  `json:"port" yaml:"port" validate:"port|eq=0"`
		Colorful bool `json:"colorful" yaml:"colorful"`
		Pprof    bool `json:"pprof" yaml:"pprof"`

		SSL        SslConfig        `json:"ssl" yaml:"ssl" validate:"dive"`
		Middleware MiddlewareConfig `json:"middleware" yaml:"middleware" validate:"dive"`
	}

	SslConfig struct {
		// Port
		/*
			0（默认）: 不使用 https port
		*/
		Port     int    `json:"port" yaml:"port" validate:"port|eq=0"`
		CertFile string `json:"certFile" yaml:"certFile"`
		KeyFile  string `json:"keyFile" yaml:"keyFile"`
	}

	MiddlewareConfig struct {
		BodyLimit     int64      `json:"bodyLimit" yaml:"bodyLimit"`
		Gzip          bool       `json:"gzip" yaml:"gzip"`
		XFrameOptions string     `json:"xFrameOptions" yaml:"xFrameOptions" validate:"omitempty,lowercase,oneof=deny sameorigin|startswith=allow-from "`
		Cors          CorsConfig `json:"cors" yaml:"cors" validate:"dive"`
		//Referer       []*refererKit.RefererVerifierBuilder `json:"referer" yaml:"referer"`
	}

	// CorsConfig cors（跨源资源共享）的配置
	CorsConfig struct {
		Access  bool     `json:"access" yaml:"access"`
		Origins []string `json:"origins" yaml:"origins" validate:"dive,required"`
	}
)
