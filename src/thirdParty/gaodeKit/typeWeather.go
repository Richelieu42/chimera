package gaodeKit

type (
	WeatherResponse struct {
		BaseResponse

		Lives     []*Live     `json:"lives"`
		Forecasts []*Forecast `json:"forecasts"`
	}

	Live struct {
		// Province 省份名
		Province string `json:"province"`
		// City 城市（区域）名
		City string `json:"city"`
		// Adcode 城市（区域）编码
		Adcode string `json:"adcode"`
		// ReportTime 数据发布的时间
		ReportTime string `json:"reporttime"`

		// Weather 天气现象（汉字描述）
		Weather string `json:"weather"`

		// Temperature 实时气温，单位：摄氏度
		Temperature string `json:"temperature"`

		// WindPower 风力级别，单位：级
		WindPower string `json:"windpower"`

		// Humidity 空气湿度
		Humidity string `json:"humidity"`
	}

	Forecast struct {
		// Province 省份名
		Province string `json:"province"`
		// City 城市（区域）名
		City string `json:"city"`
		// Adcode 城市（区域）编码
		Adcode string `json:"adcode"`
		// ReportTime 数据发布的时间
		ReportTime string `json:"reporttime"`

		// Casts 预报数据list结构，元素cast,按顺序为当天、第二天、第三天的预报数据
		Casts []*Cast `json:"casts"`
	}

	Cast struct {
		// Date 日期
		Date string `json:"date"`

		// Week 星期几
		Week string `json:"week"`

		// DayWeather 白天天气现象
		DayWeather string `json:"dayweather"`

		// NightWeather 晚上天气现象
		NightWeather string `json:"nightweather"`

		// DayTemp 白天温度
		DayTemp string `json:"daytemp"`

		// NightTemp 晚上温度
		NightTemp string `json:"nighttemp"`
	}
)
