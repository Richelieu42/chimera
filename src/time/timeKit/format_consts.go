package timeKit

import "time"

const (
	FormatCommon         TimeFormat = "2006-01-02T15:04:05.000"
	FormatCommonWithoutT TimeFormat = "2006-01-02 15:04:05.000"

	FormatRFC3339         TimeFormat = time.RFC3339 // "2006-01-02T15:04:05Z07:00"
	FormatRFC3339WithoutT TimeFormat = "2006-01-02 15:04:05Z07:00"
	FormatRFC3339Nano     TimeFormat = time.RFC3339Nano // "2006-01-02T15:04:05.999999999Z07:00"

	FormatEntire         TimeFormat = "2006-01-02T15:04:05.000Z07:00 MST"
	FormatEntireWithoutT TimeFormat = "2006-01-02 15:04:05.000Z07:00 MST"

	FormatRFC1123  TimeFormat = time.RFC1123  // "Mon, 02 Jan 2006 15:04:05 MST"
	FormatRFC1123Z TimeFormat = time.RFC1123Z // "Mon, 02 Jan 2006 15:04:05 -0700"

	FormatDateOnly TimeFormat = time.DateOnly // "2006-01-02"
	FormatTimeOnly TimeFormat = time.TimeOnly // "15:04:05"
	FormatDateTime TimeFormat = time.DateTime // "2006-01-02 15:04:05"

	// FormatFileName 用于作为文件名（或目录名）的一部分
	/*
		PS:
		(1) 不能使用 "2006-01-02T15-04-05-000".
		(2) Windows OS，文件名不支持: \ / : * ? " < > |
	*/
	FormatFileName TimeFormat = "2006-01-02T15.04.05.000"

	FormatB TimeFormat = "2006-01-02 3:04:05.000 PM Mon Jan"
	FormatC TimeFormat = "3:04:05.000 PM Mon Jan"
)
