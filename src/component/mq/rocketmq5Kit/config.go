package rocketmq5Kit

import "github.com/apache/rocketmq-clients/golang/v5/credentials"

type (
	Config struct {
		Endpoints []string `json:"endpoints" yaml:"endpoints" validate:"required,gte=1,unique,dive,hostname_port"`

		Credentials *credentials.SessionCredentials `json:"credentials" yaml:"credentials"`
	}
)
