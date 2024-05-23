package jsonFieldKit

import "github.com/tidwall/gjson"

func GetStringSliceField(jsonData []byte, path string) (s []string) {
	result := gjson.GetBytes(jsonData, path)

	for _, item := range result.Array() {
		s = append(s, item.String())
	}
	return
}

func GetStringSliceFieldFromString(jsonStr, path string) (s []string) {
	result := gjson.Get(jsonStr, path)

	for _, item := range result.Array() {
		s = append(s, item.String())
	}
	return
}

func GetInt64SliceField(jsonData []byte, path string) (s []int64) {
	result := gjson.GetBytes(jsonData, path)

	for _, item := range result.Array() {
		s = append(s, item.Int())
	}
	return
}

func GetInt64SliceFieldFromString(jsonStr, path string) (s []int64) {
	result := gjson.Get(jsonStr, path)

	for _, item := range result.Array() {
		s = append(s, item.Int())
	}
	return
}

func GetFloat64SliceField(jsonData []byte, path string) (s []float64) {
	result := gjson.GetBytes(jsonData, path)

	for _, item := range result.Array() {
		s = append(s, item.Float())
	}
	return
}

func GetFloat64SliceFieldFromString(jsonStr, path string) (s []float64) {
	result := gjson.Get(jsonStr, path)

	for _, item := range result.Array() {
		s = append(s, item.Float())
	}
	return
}

func GetBoolSliceField(jsonData []byte, path string) (s []bool) {
	result := gjson.GetBytes(jsonData, path)

	for _, item := range result.Array() {
		s = append(s, item.Bool())
	}
	return
}

func GetBoolSliceFieldFromString(jsonStr, path string) (s []bool) {
	result := gjson.Get(jsonStr, path)

	for _, item := range result.Array() {
		s = append(s, item.Bool())
	}
	return
}
