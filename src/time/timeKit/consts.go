package timeKit

import (
	"time"
)

const (
	Nanosecond  time.Duration = time.Nanosecond
	Microsecond               = time.Microsecond
	Millisecond               = time.Millisecond
	Second                    = time.Second
	Minute                    = time.Minute
	Hour                      = time.Hour
	Day                       = 24 * time.Hour
	Week                      = 7 * Day

	HalfHour = time.Minute * 30
)
