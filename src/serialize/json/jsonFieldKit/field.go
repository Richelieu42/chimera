package jsonFieldKit

import "github.com/tidwall/gjson"

// GetField
/*
@param path 如果是嵌套的内部字段，可以通过'.'组合

字段类型判断:
	Result.Type
	Result.IsBool()
	Result.IsArray()
	Result.IsObject()
*/
var GetField func(jsonData []byte, path string) gjson.Result = gjson.GetBytes

var GetFieldFromString func(json, path string) gjson.Result = gjson.Get

func GetStringField(jsonData []byte, path string) string {
	result := GetField(jsonData, path)
	return result.String()
}

func GetInt64Field(jsonData []byte, path string) int64 {
	result := GetField(jsonData, path)
	return result.Int()
}

func GetFloat64Field(jsonData []byte, path string) float64 {
	result := GetField(jsonData, path)
	return result.Float()
}

func GetBoolField(jsonData []byte, path string) bool {
	result := GetField(jsonData, path)
	return result.Bool()
}

func GetStringFieldFromString(jsonStr, path string) string {
	result := GetFieldFromString(jsonStr, path)
	return result.String()
}

func GetInt64FieldFromString(jsonStr, path string) int64 {
	result := GetFieldFromString(jsonStr, path)
	return result.Int()
}

func GetFloat64FieldFromString(jsonStr, path string) float64 {
	result := GetFieldFromString(jsonStr, path)
	return result.Float()
}

func GetBoolFieldFromString(jsonStr, path string) bool {
	result := GetFieldFromString(jsonStr, path)
	return result.Bool()
}
