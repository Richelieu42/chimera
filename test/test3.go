package main

import (
	"bufio"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v3/src/core/memoryKit"
	"github.com/richelieu-yang/chimera/v3/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"io"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Millisecond * 200)
			stats := memoryKit.GetProgramMemoryStats()

			fmt.Println(dataSizeKit.ToReadableIecString(float64(stats.Alloc)))
			fmt.Println(dataSizeKit.ToReadableIecString(float64(stats.TotalAlloc)))
			fmt.Println(dataSizeKit.ToReadableIecString(float64(stats.Sys)))
			fmt.Println("---")
		}
	}()

	path := "/Users/richelieu/Documents/ino.7z"

	f, err := fileKit.Open(path)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	var reader io.Reader
	//reader = f
	reader = bufio.NewReader(f)
	_, err = ioKit.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Since(start))
}
