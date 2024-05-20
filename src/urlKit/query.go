package urlKit

import (
	"net/url"
)

// ParseQuery
/*
e.g.
	u, err := urlKit.Verify("http://localhost:8080/go?a=123&b=456")
	if err != nil {
		panic(err)
	}
	fmt.Println(u.RawQuery) // a=123&b=456

	m, err := urlKit.ParseQuery(u.RawQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println(m) 			// map[a:[123] b:[456]]
*/
var ParseQuery func(query string) (url.Values, error) = url.ParseQuery

// ToEscapedQueryString
/*
@return 处理过的query string（可直接放到url中）
*/
func ToEscapedQueryString[T ~map[string][]string](m T) string {
	var values = url.Values(m)
	return values.Encode()
}

// AddQueryParamsToValues
/*
@param queryParams 切片中的字符串应当是未处理（编码）过的
@return 必定不为nil
*/
func AddQueryParamsToValues(values url.Values, queryParams map[string][]string) url.Values {
	if values == nil {
		values = make(map[string][]string)
	}

	for k, s := range queryParams {
		for _, v := range s {
			values.Add(k, v)
		}
	}
	return values
}

// OverrideRawQuery 覆盖请求的query.
/*
PS: 可能会修改 req.URL.RawQuery.
*/
func OverrideRawQuery(u *url.URL, queryParams map[string][]string) {
	values := AddQueryParamsToValues(nil, queryParams)
	u.RawQuery = values.Encode()
}

// AddToRawQuery 增添请求的query.
/*
PS: 可能会修改 req.URL.RawQuery.

@param queryParams 切片中的字符串应当是未处理（编码）过的
*/
func AddToRawQuery(u *url.URL, queryParams map[string][]string) {
	values := u.Query()
	values = AddQueryParamsToValues(values, queryParams)
	u.RawQuery = values.Encode()
}

func AddQueryParamsToUrl(rawURL string, queryParams map[string][]string) (string, error) {
	u, err := ParseRequestURI(rawURL)
	if err != nil {
		return "", err
	}

	/*
		!!!: 不要只使用 URL.String() ，原因: 该方法内部直接使用了 RawQuery 属性（满足条件的话），导致如果 RawQuery 中包含未处理字符（比如中文），返回值中还是会包含未处理字符
		TODO: 后续看官方是否会修改 URL.String() 中对query的处理.
	*/
	AddToRawQuery(u, queryParams)

	return u.String(), nil
}

//// ToQueryString Deprecated: use url.Values instead.
///*
//@param m 会对值进行 编码 操作
//@return 可能为""
//
//e.g.
//	(nil) => ""
//e.g.1
//	m := map[string]string{
//		"a": "test",
//		"b": "测试",
//	}
//	fmt.Println(urlKit.ToQueryString(m)) // a=test&b=%E6%B5%8B%E8%AF%95
//*/
//func ToQueryString(queryParams map[string]string) string {
//	var str string
//
//	for k, v := range queryParams {
//		// PS: k和v都有可能是""
//		if strKit.IsNotEmpty(str) {
//			str += "&"
//		}
//		if strKit.IsNotEmpty(v) {
//			str += k + "=" + EncodeURIComponent(v)
//		} else {
//			str += k
//		}
//	}
//	return str
//}
