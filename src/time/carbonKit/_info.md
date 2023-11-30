## Carbon
github(4k Star):
    https://github.com/golang-module/carbon/blob/master/README.cn.md
「Go工具箱」推荐一个轻量级、语义化的时间处理库：carbon
    https://mp.weixin.qq.com/s/92O1SHs4tw1FMcDeQoSfkA

这是一个轻量级的、易于使用的、语义智能的日期时间库，适用于Go开发者。

#### Carbon 和 time.Time 互转
// 将 time.Time 转换成 Carbon
carbon.CreateFromStdTime(time.Now())

// 将 Carbon 转换成 time.Time
carbon.Now().ToStdTime()
