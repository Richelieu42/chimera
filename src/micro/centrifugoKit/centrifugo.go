package centrifugoKit

import "github.com/centrifugal/gocent/v3"

var (
	NewClient func(c gocent.Config) *gocent.Client = gocent.New
)

func a() {

}
