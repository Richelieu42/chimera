// Package confKit
/*
支持的配置文件格式（详见viper.go）:
"yaml", "yml", "json", "toml", "hcl", "tfvars",
"dotenv", "env",
"properties", "props", "prop",
"ini"
*/
package confKit

import (
	"github.com/mitchellh/mapstructure"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/ioKit"
	"github.com/richelieu42/chimera/src/core/pointerKit"
	"github.com/spf13/viper"
)

// ReadAs 读取配置文本内容，并反序列化.
/*
Deprecated: 推荐使用 Load || MustLoad.

@param configType 	配置文件的类型（不区分大小写，详见viper.go）："yaml", "yml", "json", "toml", "hcl", "tfvars", "dotenv", "env", "properties", "props", "prop", "ini"
@param defaultMap 	默认值，可以为nil
@param ptr			指针，且不能为nil
*/
func ReadAs(data []byte, configType string, defaultMap map[string]interface{}, ptr interface{}) error {
	if err := checkParamPtr(ptr); err != nil {
		return err
	}

	v := viper.New()
	for key, value := range defaultMap {
		v.SetDefault(key, value)
	}
	v.SetConfigType(configType)
	if err := v.ReadConfig(ioKit.NewReader(data)); err != nil {
		return err
	}
	return v.Unmarshal(ptr, func(dc *mapstructure.DecoderConfig) {
		// 如果指针ptr对应的类型是个子类的话，需要如此进行设置，否则父类的属性都会是对应类型的零值！
		dc.Squash = true
	})
}

// ReadFileAs 读取配置文件，并反序列化.
/*
Deprecated: 推荐使用 Load || MustLoad.

PS:
(1) 配置文件 和 defaultMap 中，key首字母的大小写无所谓，都支持；
(2) 支持配置文件的格式：JSON, TOML, HCL, .env, .yaml, .properties.

@param filePath		配置文件的路径（绝对路径 和 相对路径 都支持），内部会判断文件是否存在
@param defaultMap 	（可以为nil） 默认值；key如果有多层的话，用"."分隔，e.g. "WoService.LowerLimit"
@param ptr			（不能为nil） 指针
*/
func ReadFileAs(filePath string, defaultMap map[string]interface{}, ptr interface{}) error {
	if err := checkParamPtr(ptr); err != nil {
		return err
	}

	v, err := readFile(filePath, defaultMap)
	if err != nil {
		return err
	}
	return v.Unmarshal(ptr, func(dc *mapstructure.DecoderConfig) {
		// 如果指针ptr对应的类型是个子类的话，需要如此进行设置，否则父类的属性都会是对应类型的零值！
		dc.Squash = true
	})
}

// checkParamPtr 检查传参ptr：要求是指针且不为nil
func checkParamPtr(ptr interface{}) error {
	if ptr == nil {
		return errorKit.Simple("ptr is nil")
	}
	if !pointerKit.IsPointer(ptr) {
		return errorKit.Simple("ptr isn't a pointer")
	}
	return nil
}

func readFile(filePath string, defaultMap map[string]interface{}) (*viper.Viper, error) {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return nil, err
	}

	v := viper.New()
	v.SetConfigFile(filePath)

	// 设置默认值
	for key, value := range defaultMap {
		v.SetDefault(key, value)
	}

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}
