package timeKit

import (
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	var err error

	name := "GMT"
	GMT, err = time.LoadLocation(name)
	if err != nil {
		logrus.WithError(err).Fatalf("Fail to load location with name(%s).", name)
	}
}
