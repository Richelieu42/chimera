package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/time/timeKit"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(timeKit.Format(t, timeKit.FormatNetwork))
	fmt.Println(timeKit.Format(t.In(time.UTC), timeKit.FormatNetwork))
	fmt.Println(timeKit.Format(t.In(timeKit.GMT), timeKit.FormatNetwork))

	loc := time.FixedZone("GMT", 0)
	fmt.Println(timeKit.Format(t.In(loc), timeKit.FormatNetwork))
}
