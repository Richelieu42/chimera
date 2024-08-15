package slbKit

import "time"

type (
	Status string
)

const (
	StatusInitialized Status = "initialized"

	StatusStarted Status = "started"

	StatusDisposed Status = "disposed"
)

const (
	// HealthCheckInterval 健康检查的周期
	HealthCheckInterval = time.Second * 10

	// HealthCheckTimeout 健康检查的超时时间
	HealthCheckTimeout = time.Second * 3
)
