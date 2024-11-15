package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := `1、开局就迎来破灭是不是搞错了什么

“我这是……穿越了？”
2、难度还要增加是吧

“系统？在吗？”`

	re, err := regexp.Compile(`(\d+)、(.+)`)
	if err != nil {
		panic(err)
	}
	content1 := re.ReplaceAllStringFunc(content, func(str string) string {
		s := re.FindStringSubmatch(str)
		s1 := fmt.Sprintf("第%s章、%s", s[1], s[2])
		return s1
	})
	fmt.Println(content1)
	/*
		第1章、开局就迎来破灭是不是搞错了什么

		“我这是……穿越了？”
		第2章、难度还要增加是吧

		“系统？在吗？”
	*/
}
