package centrifugoKit

import "github.com/centrifugal/gocent/v3"

var (
	NerClient func(c gocent.Config) *gocent.Client = gocent.New
)
