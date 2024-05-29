package validateKit

// Required 必填，非零值（zero value）
/*
	e.g.
		fmt.Println(validateKit.Required(nil)) 		// Key: '' Error:Var validation for '' failed on the 'required' tag

		fmt.Println(validateKit.Required(""))    	// Key: '' Error:Var validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required("aaa")) 	// <nil>

		fmt.Println(validateKit.Required(0)) 		// Key: '' Error:Var validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required(1)) 		// <nil>

		fmt.Println(validateKit.Required(false)) 	// Key: '' Error:Var validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required(true))  	// <nil>
*/
func Required(field interface{}) error {
	return Var(field, "required")
}

// IP
/*
	e.g.
		fmt.Println(validateKit.IP(""))          // Key: '' Error:Var validation for '' failed on the 'ip' tag
		fmt.Println(validateKit.IP("127.0.0.1")) // <nil>
		fmt.Println(validateKit.IP("127.001"))   // Key: '' Error:Var validation for '' failed on the 'ip' tag
*/
func IP(field interface{}) error {
	return Var(field, "ip")
}

func IPv4(field interface{}) error {
	return Var(field, "ipv4")
}

func Email(field interface{}) error {
	return Var(field, "email")
}

// HttpUrl
/*
	PS: 要以 "http://" 或 "https://" 开头.

	e.g.
		fmt.Println(validateKit.HttpUrl(""))                                           // Key: '' Error:Var validation for '' failed on the 'http_url' tag
		fmt.Println(validateKit.HttpUrl("https://github.com/go-playground/validator")) // <nil>
		fmt.Println(validateKit.HttpUrl("http://github.com/go-playground/validator"))  // <nil>
		fmt.Println(validateKit.HttpUrl("ftp://github.com/go-playground/validator"))   // Key: '' Error:Var validation for '' failed on the 'http_url' tag
*/
func HttpUrl(field interface{}) error {
	return Var(field, "http_url")
}

// Json 字符串值是否为有效的JSON.
/*
	e.g.
		fmt.Println(validateKit.Json(""))   // Key: '' Error:Var validation for '' failed on the 'json' tag
		fmt.Println(validateKit.Json("[]")) // <nil>
		fmt.Println(validateKit.Json("{}")) // <nil>
		fmt.Println(validateKit.Json("[}")) // Key: '' Error:Var validation for '' failed on the 'json' tag
*/
func Json(field interface{}) error {
	return Var(field, "json")
}

// File 字符串值是否包含有效的文件路径，以及该文件是否存在于计算机上.
/*
	PS: 传参对应的应当是"文件"，是"目录"的话会返回error.

	e.g.
		fmt.Println(validateKit.File("")) // Key: '' Error:Var validation for '' failed on the 'file' tag

		// 目录存在
		fmt.Println(validateKit.File("_chimera-lib"))                                         // Key: '' Error:Var validation for '' failed on the 'file' tag
		fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/_chimera-lib")) // Key: '' Error:Var validation for '' failed on the 'file' tag
		// 文件存在
		fmt.Println(validateKit.File("_chimera-lib/config.yaml"))                                         // <nil>
		fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/_chimera-lib/config.yaml")) // <nil>
		// 文件不存在
		fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/_chimera-lib/config111.yaml")) // Key: '' Error:Var validation for '' failed on the 'file' tag
		// 无效的文件路径
		fmt.Println(validateKit.File("chimera-lib\\config.yaml")) // Key: '' Error:Var validation for '' failed on the 'file' tag
*/
func File(field interface{}) error {
	return Var(field, "file")
}

func Port(field interface{}) error {
	return Var(field, "gt=0,lte=65535")
}

// Hostname
/*
通过的field: "www.yozo.com"、"10.0.9.141"
*/
func Hostname(field interface{}) error {
	return Var(field, "hostname|ipv4")
}

// Host
/*
通过的field: "www.yozo.com:8888"、"10.0.9.141:80"
不能通过的field: "10.0.9.141:0"、"10.0.9.141:-1"
*/
func Host(field interface{}) error {
	return Var(field, "hostname_port")
}
