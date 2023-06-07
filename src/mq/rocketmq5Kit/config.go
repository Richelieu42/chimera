package rocketmq5Kit

import "github.com/apache/rocketmq-clients/golang/v5/credentials"

type (
	Config struct {
		Endpoints      []string                        `json:"endpoints"`
		Credentials    *credentials.SessionCredentials `json:"credentials,optional"`
		ValidatedTopic string                          `json:"validatedTopic,optional"`
		ClientLogPath  string                          `json:"clientLogPath,optional"`
	}
)
