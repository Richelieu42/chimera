package req111Kit

func Post(url string, queryParams map[string][]string, body interface{}) (status int, data []byte, err error) {
	return GetDefaultClient().Post(url, queryParams, body)
}
