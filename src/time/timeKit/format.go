package timeKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"time"
)

type (
	TimeFormat string
)

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

// FormatCurrent 格式化 当前时间 为 字符串.
/*
@param format 可以为""

e.g.
	("2006-01-02T15:04:05.000") => "2023-08-17T16:05:14.985"
*/
func FormatCurrent[F ~string](format F) string {
	if strKit.IsEmpty(string(format)) {
		format = F(FormatCommon)
	}

	return time.Now().Format(string(format))
}

// Format time.Time => string
/*
@param t		不用担心t为nil的情况，详见下面的说明
@param formats 	不传的话用默认值；传多个（包括一个）的话用第一个

一个方法如果接受类型为time.Time的参数，那么不用考虑该参数为nil的情况，因为：
（1）time.Time类型变量的零值不为nil；
（2）调用时，该参数位置不能直接传参nil（IDE报错：Cannot use 'nil' as the type time.Time）；
（3）time.Time类型变量不能被赋值为nil（IDE报错：Cannot use 'nil' as the type time.Time）.

e.g.
	str := timeKit.Format(time.Now(), timeKit.FormatCommon)
	fmt.Println(str)	// 2023-08-14T17:10:17.057
*/
func Format[F ~string](t time.Time, format F) string {
	return t.Format(string(format))
}

// FormatDuration time.Duration => string
/*
e.g.
	d := time.Minute*63 + time.Second*15
	fmt.Println(timeKit.FormatDuration(d)) // 1h3m15s
*/
func FormatDuration(d time.Duration) string {
	return d.String()
	//return fmt.Sprintf("%s", d)
}
