package jsonFieldKit

import (
	"fmt"
	"github.com/tidwall/gjson"
	"testing"
)

func TestGetStringField(t *testing.T) {
	jsonStr := `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	fmt.Println(GetStringField([]byte(jsonStr), "name.last")) // Prichard
}

func TestGetField(t *testing.T) {
	{
		json := `{"status":"1","info":"OK","infocode":"10000","province":[],"city":[],"adcode":[],"rectangle":[]}`
		f := GetField([]byte(json), "province")
		fmt.Println(f)
		fmt.Println(f.Type)
		fmt.Println("IsString", f.Type == gjson.String)
		fmt.Println("IsArray", f.IsArray())
	}
	fmt.Println("------")
	{
		json := `{"status":"1","info":"OK","infocode":"10000","province":"局域网","city":[],"adcode":[],"rectangle":[]}`
		f := GetField([]byte(json), "province")
		fmt.Println(f)
		fmt.Println(f.Type)
		fmt.Println("IsString", f.Type == gjson.String)
		fmt.Println("IsArray", f.IsArray())
	}
	/*
		[]
		JSON
		IsString false
		IsArray true
		------
		局域网
		String
		IsString true
		IsArray false
	*/
}
