package timeKit

import (
	"time"
)

type (
	TimeFormat string
)

// FormatCurrent 格式化 当前时间 为 字符串.
/*
@param formats 可以为""

e.g.
	("2006-01-02T15:04:05.000") => "2023-08-17T16:05:14.985"
*/
func FormatCurrent[F ~string](f F) string {
	var format = string(f)
	if format == "" {
		format = string(FormatEntire)
	}

	return time.Now().Format(format)
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
