package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"io"
	"os"
)

func main() {
	tmp := &os.File{}
	//tmp := &bufio.Writer{}
	//tmp := &bufio.ReadWriter{}

	var _ io.Reader = tmp
	//var _ io.ReaderAt = tmp
	//var _ io.ReaderFrom = tmp

	var _ io.Writer = tmp
	//var _ io.WriterAt = tmp
	//var _ io.WriterTo = tmp

	var _ io.Seeker = tmp

	var _ io.Closer = tmp
}
