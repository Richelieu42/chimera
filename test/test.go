package main

import (
	"github.com/richelieu-yang/chimera/v3/src/core/osKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
)

func main() {
	sliceKit.RemoveByIndex()

	console.Infof("os: %s", osKit.OS)
	console.Infof("arch: %s", osKit.ARCH)
	console.Infof("json library: %s", jsonKit.GetLibrary())

	json, err := jsonKit.MarshalToString(map[string]interface{}{
		"a": true,
		"0": []interface{}{true, false, 0, "ccc"},
	})
	if err != nil {
		panic(err)
	}
	console.Info(json)
}
