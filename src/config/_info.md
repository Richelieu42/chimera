## Go几种读取配置文件的方式

https://mp.weixin.qq.com/s/LCbA4r8euBzHIUJjGBBxFg

#### spf13/viper

    https://github.com/spf13/viper

#### jinzhu/configor（最后更新时间: 2020/11/18）

    https://github.com/jinzhu/configor

## spf13/viper VS go-zero conf VS json-iterator/go

json字符串 反序列化为 结构体实例（某些值应当为int类型，但json中的却是string类型）
(a) spf13/viper:        反序列化成功（但spf13/viper不支持通过json tag设置默认值）.    
(b) go-zero conf:       反序列化直接返回error.
(c) json-iterator/go:   反序列化直接返回error.
