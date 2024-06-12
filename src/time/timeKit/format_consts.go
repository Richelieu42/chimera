package timeKit

import "time"

const (
	FormatRFC3339     TimeFormat = time.RFC3339
	FormatRFC3339Nano TimeFormat = time.RFC3339Nano

	// FormatNetwork 网络的格式
	FormatNetwork  TimeFormat = FormatRFC1123
	FormatRFC1123  TimeFormat = time.RFC1123
	FormatRFC1123Z TimeFormat = time.RFC1123Z

	// FormatDate 日期格式（年月日）
	FormatDate TimeFormat = "2006-01-02"
	// FormatDateTime 日期时间格式（年月日时分秒）
	FormatDateTime TimeFormat = time.DateTime

	// FormatFileName 用于作为文件名（或目录名）的一部分
	/*
		PS:
		(1) 不能使用 "2006-01-02T15-04-05-000".
		(2) Windows OS，文件名不支持: \ / : * ? " < > |
	*/
	FormatFileName TimeFormat = "2006-01-02T15.04.05.000"

	FormatCommon      TimeFormat = "2006-01-02 15:04:05.000"
	FormatCommonWithT TimeFormat = "2006-01-02T15:04:05.000"

	FormatEntire      TimeFormat = "2006-01-02 15:04:05.000Z07:00 MST"
	FormatEntireWithT TimeFormat = "2006-01-02T15:04:05.000Z07:00 MST"

	FormatB TimeFormat = "2006-01-02 3:04:05.000 PM Mon Jan"
	FormatC TimeFormat = "3:04:05.000 PM Mon Jan"
)
