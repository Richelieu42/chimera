package req111Kit

func Get(url string, queryParams map[string][]string) (code int, data []byte, err error) {
	return GetDefaultClient().Get(url, queryParams)
}
