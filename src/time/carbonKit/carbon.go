package carbonKit

import "github.com/dromara/carbon/v2"

var (
	// Yesterday 昨天.
	Yesterday func(timezone ...string) carbon.Carbon = carbon.Yesterday

	// Now 当前.
	Now func(timezone ...string) carbon.Carbon = carbon.Now

	// Tomorrow 明天.
	Tomorrow func(timezone ...string) carbon.Carbon = carbon.Tomorrow
)
