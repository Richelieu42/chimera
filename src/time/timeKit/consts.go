package timeKit

import (
	"time"
)

type (
	TimeFormat string
)

const (
	Nanosecond  time.Duration = time.Nanosecond
	Microsecond               = time.Microsecond
	Millisecond               = time.Millisecond
	Second                    = time.Second
	Minute                    = time.Minute
	Hour                      = time.Hour
	Day                       = 24 * time.Hour
	Week                      = 7 * Day

	HalfHour = time.Minute * 30
)

/*
UTC: 世界协调时间(Coordinated Universal Time)
GMT: 格林威治时间(UTC+0)
CST: 中国标准时间(UTC+8)
MST: 北美山区标准时间(UTC-7)
*/
const (
	FormatDate TimeFormat = "2006-01-02"

	// FormatDefault 参考 time/format.go
	FormatDefault TimeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

	// FormatFileName 用于作为文件名（或目录名）的一部分
	/*
		PS:
		(1) 不能使用 "2006-01-02T15-04-05-000".
		(2) Windows OS，文件名不支持: \ / : * ? " < > |
	*/
	FormatFileName TimeFormat = "2006-01-02T15.04.05.000"

	// FormatCommon 常规的格式
	FormatCommon TimeFormat = "2006-01-02T15:04:05.000"
	// FormatCommon1 常规的格式1
	FormatCommon1 TimeFormat = "2006-01-02 15:04:05.000"

	// FormatEntire 完整的格式
	FormatEntire  TimeFormat = "2006-01-02 15:04:05.000Z07:00 MST"
	FormatEntire1 TimeFormat = "2006-01-02T15:04:05.000Z07:00 MST"
	FormatEntire2 TimeFormat = "2006-01-02T15:04:05.000Z07:00"

	// FormatNetwork 网络的格式
	FormatNetwork TimeFormat = "Mon, 02 Jan 2006 15:04:05 MST"

	// FormatRFC1123 网络的格式
	FormatRFC1123 TimeFormat = time.RFC1123

	FormatA TimeFormat = "2006-01-02 15:04:05"
	FormatB TimeFormat = "2006-01-02 3:04:05.000 PM Mon Jan"
	FormatC TimeFormat = "3:04:05.000 PM Mon Jan"
)
