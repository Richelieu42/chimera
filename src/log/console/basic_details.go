package console

import (
	"github.com/richelieu-yang/chimera/v3/src/log/commonLogKit"
)

func PrintBasicDetails() {
	l := newLogger(0).Sugar()
	defer l.Sync()

	commonLogKit.PrintBasicDetails(l)
}
