## go-zero/core/conf(v1.5.6) 支持读取的配置文件格式

* .json
* .toml
* .yaml
* .yml
* .properties（比较low）

#### 读取 .properties 格式的文件

虽然"github.com/zeromicro/go-zero/core/conf"支持读取.properties格式的文件，但目前还不能反序列化，比较low，建议此种情况下还是使用viper.
