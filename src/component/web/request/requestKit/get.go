package requestKit

import "github.com/imroc/req/v3"

// SimpleGet
/*
TODO: 会不会重试？
*/
func SimpleGet(url string) (statusCode int, data []byte, err error) {
	return Get(url, nil)
}

// Get
/*
TODO: 会不会重试？
*/
func Get(url string, queryParams map[string]interface{}) (statusCode int, data []byte, err error) {
	r := NewRequest(nil)
	var resp *req.Response
	resp, err = r.SetQueryParamsAnyType(queryParams).Get(url)
	if err != nil {
		return
	}
	// 不需要手动关闭 resp
	//defer resp.Body.Close()

	statusCode = resp.StatusCode
	data = resp.Bytes()
	return
}
