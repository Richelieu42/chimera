package main

import (
	"bufio"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"io"
	"time"
)

func main() {
	//go func() {
	//	for {
	//		time.Sleep(time.Millisecond * 200)
	//		stats := memoryKit.GetProgramMemoryStats()
	//
	//		fmt.Println(dataSizeKit.ToReadableIecString(float64(stats.Alloc)))
	//		fmt.Println(dataSizeKit.ToReadableIecString(float64(stats.TotalAlloc)))
	//		fmt.Println(dataSizeKit.ToReadableIecString(float64(stats.Sys)))
	//		fmt.Println("---")
	//	}
	//}()

	start := time.Now()

	for i := 0; i < 10; i++ {
		path := "/Users/richelieu/Documents/ino.7z"
		f, err := fileKit.Open(path)
		if err != nil {
			panic(err)
		}
		var reader io.Reader
		//reader = f
		reader = bufio.NewReader(f)
		_, err = ioKit.ReadAll(reader)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(time.Since(start))
}
