package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile(`(?i)^CaSe.*`)

	fmt.Println(r.MatchString("case"))  // true
	fmt.Println(r.MatchString("CASE"))  // true
	fmt.Println(r.MatchString("CAse"))  // true
	fmt.Println(r.MatchString("1CAse")) // false
}

//// VerifyReferer 验证 referer
///*
//@return 验证是否成功 + 验证失败的原因
//*/
//func VerifyReferer(referer string, none bool, blocked bool, res []*regexp.Regexp) (bool, string) {
//	if strKit.IsEmpty(referer) {
//		return none, "none"
//	}
//
//	var prefix string
//	if strKit.StartWith(referer, "http://") {
//		prefix = "http://"
//	} else if strKit.StartWith(referer, "https://") {
//		prefix = "https://"
//	} else {
//		return blocked, "blocked"
//	}
//
//	referer = strKit.RemovePrefixIfExist(referer, prefix)
//	referer = strKit.SubBeforeString(referer, "/")
//	// 忽略端口号（有的话）
//	referer = strKit.SubBeforeString(referer, ":")
//
//	for _, re := range res {
//		if re.MatchString(referer) {
//			return true, ""
//		}
//	}
//	return false, "no match found"
//
//	//i := strKit.Index(referer, "/")
//	//if i != -1 {
//	//	referer = strKit.SubBefore(referer, i)
//	//}
//	//i = strKit.Index(referer, ":")
//	//if i != -1 {
//	//	referer = strKit.SubBefore(referer, i)
//	//}
//
//	return false, ""
//}
