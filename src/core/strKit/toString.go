package strKit

import (
	"github.com/spf13/cast"
	"strconv"
)

var (
	ToString func(i interface{}) string = cast.ToString

	ToStringE func(i interface{}) (string, error) = cast.ToStringE

	// IntToString int（10进制） => string
	/*
		PS: strconv.Itoa(i) <=> FormatInt(int64(i), 10)
	*/
	IntToString func(i int) string = strconv.Itoa

	// IntToStringWithBase int（各种进制） => string
	/*
		@param base 传参i的进制数
	*/
	IntToStringWithBase func(i int64, base int) string = strconv.FormatInt
)
