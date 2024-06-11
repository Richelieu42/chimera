package jsonKit

/*
	Richelieu: "github.com/json-iterator/go" 已经很久不更新了，还是用 官方库 或 sonic.
*/

////go:build (linux || windows || darwin) && !(sonic && avx && go1.17 && amd64)
//
//package jsonKit
//
//import (
//	jsoniter "github.com/json-iterator/go"
//)
//
//func init() {
//	library = "json-iterator/go"
//	defaultApi = jsoniter.ConfigDefault
//	stdApi = jsoniter.ConfigCompatibleWithStandardLibrary
//
//	testAPI()
//}
