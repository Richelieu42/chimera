// Package consts 本库的全局常量.
package consts

import (
	_ "embed"
)

const (
	// ProjectName 本库（项目）的名字
	ProjectName = "chimera"

	// LowerProjectName 全小写
	LowerProjectName = "chimera"

	// UpperProjectName 全大写
	UpperProjectName = "CHIMERA"
)

// Banner
/*
我把SpringBoot的banner换成了美女，老板说工作不饱和，建议安排加班...
	https://mp.weixin.qq.com/s/YJJp2zrvGfFXqEnYIvVycQ
英文ASCII艺术字，Spring Boot自定义启动Banner在线生成工具
	https://www.bootschool.net/ascii
*/
//go:embed _banner.txt
var Banner string

//func init() {
//	data, err := resources.Asset("_resources/banner.txt")
//	if err != nil {
//		logrus.WithError(err).Fatalf("[%s] Fail to get banner.", UpperProjectName)
//		return
//	}
//	Banner = string(data)
//}
