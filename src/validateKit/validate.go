package validateKit

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
	"reflect"
)

// New
/*
PS:
(1) 默认的tag name: "validate";
(2) Gin使用的tag name: "binding".

@param tagNameArgs 不传参的话，将采用默认的tag name
*/
func New(tagNameArgs ...string) *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())

	tagName := sliceKit.GetFirstItemWithDefault("", tagNameArgs...)
	if strKit.IsNotEmpty(tagName) {
		v.SetTagName(tagName)
	}

	if err := registerDefaultValidation(v); err != nil {
		panic(err)
	}

	return v
}

// registerDefaultValidation 注册默认的验证器(s).
func registerDefaultValidation(v *validator.Validate) error {
	tag := "port"
	err := v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		var field reflect.Value = fl.Field()
		return netKit.IsValidPort(field)
	})
	if err != nil {
		return err
	}

	return nil
}

// Struct 验证结构体.
/*
@param s 如果为nil，将返回error(e.g. validator: (nil *main.User))
*/
func Struct(s interface{}) error {
	if validatable, ok := s.(Validatable); ok {
		return validatable.Validate()
	}

	v := New()
	return v.Struct(s)
}

// Field 验证字段.
func Field(field interface{}, tag string) error {
	v := New()
	return v.Var(field, tag)
}

// Required 必填，非零值（zero value）
/*
	e.g.
		fmt.Println(validateKit.Required(nil)) 		// Key: '' Error:Field validation for '' failed on the 'required' tag

		fmt.Println(validateKit.Required(""))    	// Key: '' Error:Field validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required("aaa")) 	// <nil>

		fmt.Println(validateKit.Required(0)) 		// Key: '' Error:Field validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required(1)) 		// <nil>

		fmt.Println(validateKit.Required(false)) 	// Key: '' Error:Field validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required(true))  	// <nil>
*/
func Required(field interface{}) error {
	return Field(field, "required")
}

// IP
/*
	e.g.
		fmt.Println(validateKit.IP(""))          // Key: '' Error:Field validation for '' failed on the 'ip' tag
		fmt.Println(validateKit.IP("127.0.0.1")) // <nil>
		fmt.Println(validateKit.IP("127.001"))   // Key: '' Error:Field validation for '' failed on the 'ip' tag
*/
func IP(field interface{}) error {
	return Field(field, "ip")
}

func IPv4(field interface{}) error {
	return Field(field, "ipv4")
}

func Email(field interface{}) error {
	return Field(field, "email")
}

// HttpUrl
/*
	PS: 要以 "http://" 或 "https://" 开头.

	e.g.
		fmt.Println(validateKit.HttpUrl(""))                                           // Key: '' Error:Field validation for '' failed on the 'http_url' tag
		fmt.Println(validateKit.HttpUrl("https://github.com/go-playground/validator")) // <nil>
		fmt.Println(validateKit.HttpUrl("http://github.com/go-playground/validator"))  // <nil>
		fmt.Println(validateKit.HttpUrl("ftp://github.com/go-playground/validator"))   // Key: '' Error:Field validation for '' failed on the 'http_url' tag
*/
func HttpUrl(field interface{}) error {
	return Field(field, "http_url")
}

// Json 字符串值是否为有效的JSON.
/*
	e.g.
		fmt.Println(validateKit.Json(""))   // Key: '' Error:Field validation for '' failed on the 'json' tag
		fmt.Println(validateKit.Json("[]")) // <nil>
		fmt.Println(validateKit.Json("{}")) // <nil>
		fmt.Println(validateKit.Json("[}")) // Key: '' Error:Field validation for '' failed on the 'json' tag
*/
func Json(field interface{}) error {
	return Field(field, "json")
}

// File 字符串值是否包含有效的文件路径，以及该文件是否存在于计算机上.
/*
	PS: 传参对应的应当是"文件"，是"目录"的话会返回error.

	e.g.
		fmt.Println(validateKit.File("")) // Key: '' Error:Field validation for '' failed on the 'file' tag

		// 目录存在
		fmt.Println(validateKit.File("_chimera-lib"))                                         // Key: '' Error:Field validation for '' failed on the 'file' tag
		fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/_chimera-lib")) // Key: '' Error:Field validation for '' failed on the 'file' tag
		// 文件存在
		fmt.Println(validateKit.File("_chimera-lib/config.yaml"))                                         // <nil>
		fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/_chimera-lib/config.yaml")) // <nil>
		// 文件不存在
		fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/_chimera-lib/config111.yaml")) // Key: '' Error:Field validation for '' failed on the 'file' tag
		// 无效的文件路径
		fmt.Println(validateKit.File("chimera-lib\\config.yaml")) // Key: '' Error:Field validation for '' failed on the 'file' tag
*/
func File(field interface{}) error {
	return Field(field, "file")
}

func Port(field interface{}) error {
	return Field(field, "gt=0,lte=65535")
}
