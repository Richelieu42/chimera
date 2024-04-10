package timeKit

import (
	"time"
)

func init() {
	var err error
	GMT, err = time.LoadLocation("GMT")
	if err != nil {
		GMT = time.FixedZone("GMT", 0)
	}
}
