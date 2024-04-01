package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/dataSizeKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"io"
	"net/http"
)

func main() {
	io.Copy()

	var r *http.Request

	fmt.Println(dataSizeKit.ToReadableIecString(512))
	//io.ReadAll()
	//io.ReadFull()
	//
	//tp := noop.NewTracerProvider()
}
