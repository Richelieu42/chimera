package gaodeKit

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/sirupsen/logrus"
	"testing"
)

/*
&{江苏 锡山区 320205 2023-11-15 17:06:24 多云 13 ≤3 46}
&{2023-11-15 3 多云 小雨 17 9}
*/
func TestWeather(t *testing.T) {
	client, err := NewClient("b15c36bf1df4c272e92f3f1875a127f1")
	if err != nil {
		panic(err)
	}

	live, err := client.GetLive("320205")
	if err != nil {
		panic(err)
	}
	fmt.Println(live)

	todayCast, err := client.GetTodayCast("320205")
	if err != nil {
		panic(err)
	}
	fmt.Println(todayCast)
}

/*
2023-11-15 17:33:00.779+08:00 CST
+--------+----------+----------+----------+
| 地区   | 当前天气 | 白天天气 | 夜晚天气 |
+--------+----------+----------+----------+
| 无锡市 | 13(多云) | 17(多云) | 9(小雨)  |
| 惠山区 | 13(多云) | 17(多云) | 9(小雨)  |
| 滨湖区 | 13(多云) | 17(多云) | 9(小雨)  |
| 梁溪区 | 13(多云) | 17(多云) | 9(小雨)  |
| 上海市 | 14(多云) | 17(多云) | 11(小雨) |
+--------+----------+----------+----------+
*/
func TestWeather1(testing *testing.T) {
	client, err := NewClient("b15c36bf1df4c272e92f3f1875a127f1")
	if err != nil {
		panic(err)
	}

	t := table.NewWriter()
	t.AppendHeader(table.Row{"地区", "当前天气", "白天天气", "夜晚天气"})

	cities := []string{"320200", "320206", "320211", "320213", "310000"}
	for _, city := range cities {
		live, err := client.GetLive(city)
		if err != nil {
			logrus.WithError(err).WithField("city", city).Error("Fail to get live")
			continue
		}
		todayCast, err := client.GetTodayCast(city)
		if err != nil {
			logrus.WithError(err).WithField("city", city).Error("Fail to get today cast")
			continue
		}

		t.AppendRow(table.Row{live.City, fmt.Sprintf("%s(%s)",
			live.Temperature, live.Weather),
			fmt.Sprintf("%s(%s)", todayCast.DayTemp, todayCast.DayWeather),
			fmt.Sprintf("%s(%s)", todayCast.NightTemp, todayCast.NightWeather)})
	}

	fmt.Println(timeKit.FormatCurrent(timeKit.FormatEntire))
	fmt.Println(t.Render())
}
