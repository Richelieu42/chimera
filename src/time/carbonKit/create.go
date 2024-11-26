package carbonKit

import (
	"github.com/dromara/carbon/v2"
	"time"
)

var (
	// CreateFromStdTime time.Time => carbon.Carbon
	CreateFromStdTime func(tt time.Time, timezone ...string) carbon.Carbon = carbon.CreateFromStdTime

	// CreateFromDate 从给定的年、月、日创建 Carbon 实例
	CreateFromDate      func(year, month, day int, timezone ...string) carbon.Carbon              = carbon.CreateFromDate
	CreateFromDateMilli func(year, month, day, millisecond int, timezone ...string) carbon.Carbon = carbon.CreateFromDateMilli
	CreateFromDateMicro func(year, month, day, microsecond int, timezone ...string) carbon.Carbon = carbon.CreateFromDateMicro
	CreateFromDateNano  func(year, month, day, nanosecond int, timezone ...string) carbon.Carbon  = carbon.CreateFromDateNano

	// CreateFromDateTime 从给定的年、月、日、时、分、秒创建 Carbon 实例
	CreateFromDateTime      func(year, month, day, hour, minute, second int, timezone ...string) carbon.Carbon              = carbon.CreateFromDateTime
	CreateFromDateTimeMilli func(year, month, day, hour, minute, second, millisecond int, timezone ...string) carbon.Carbon = carbon.CreateFromDateTimeMilli
	CreateFromDateTimeMicro func(year, month, day, hour, minute, second, microsecond int, timezone ...string) carbon.Carbon = carbon.CreateFromDateTimeMicro
	CreateFromDateTimeNano  func(year, month, day, hour, minute, second, nanosecond int, timezone ...string) carbon.Carbon  = carbon.CreateFromDateTimeNano

	// CreateFromTimestamp 从给定的秒级时间戳创建 Carbon 实例
	CreateFromTimestamp      func(timestamp int64, timezone ...string) carbon.Carbon = carbon.CreateFromTimestamp
	CreateFromTimestampMilli func(timestamp int64, timezone ...string) carbon.Carbon = carbon.CreateFromTimestampMilli
	CreateFromTimestampMicro func(timestamp int64, timezone ...string) carbon.Carbon = carbon.CreateFromTimestampMicro
	CreateFromTimestampNano  func(timestamp int64, timezone ...string) carbon.Carbon = carbon.CreateFromTimestampNano

	// CreateFromTime 从给定的时、分、秒创建 Carbon 实例
	CreateFromTime      func(hour, minute, second int, timezone ...string) carbon.Carbon              = carbon.CreateFromTime
	CreateFromTimeMilli func(hour, minute, second, millisecond int, timezone ...string) carbon.Carbon = carbon.CreateFromTimeMilli
	CreateFromTimeMicro func(hour, minute, second, microsecond int, timezone ...string) carbon.Carbon = carbon.CreateFromTimeMicro
	CreateFromTimeNano  func(hour, minute, second, nanosecond int, timezone ...string) carbon.Carbon  = carbon.CreateFromTimeNano
)
