package baiduOcrKit

import (
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
	"gitee.com/richelieu042/go-scales/src/core/file/fileKit"
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	"gitee.com/richelieu042/go-scales/src/http/httpClientKit"
	"gitee.com/richelieu042/go-scales/src/jsonKit"
	"gitee.com/richelieu042/go-scales/src/urlKit"
)

const (
	// 固定参数
	grantType = "client_credentials"
)

// RecognizeUniversalWords 通用文字识别（标准版）
/*
文档:
https://cloud.baidu.com/doc/OCR/s/zk3h7xz52
https://ai.baidu.com/ai-doc/OCR/zk3h7xz52
*/
func RecognizeUniversalWords(imagePath string) (*Words, error) {
	// url
	token, err := getAccessToken()
	if err != nil {
		return nil, err
	}
	url := strKit.Format("%s?access_token=%s", "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic", token.AccessToken)

	// params
	base64Str, err := fileKit.ReadFileToBase64(imagePath)
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"language_type": "CHN_ENG",
		"image":         urlKit.EncodeURIComponent(base64Str),
	}

	// 发请求
	resp, err := httpClientKit.Post(url, params)
	if err != nil {
		return nil, err
	}
	m, err := jsonKit.UnmarshalToMap(resp)
	if err != nil {
		return nil, err
	}

	// 解析响应
	words, err := parseMapToWords(m)
	if err != nil {
		return nil, err
	}
	if words == nil {
		return nil, errorKit.Simple("failure response(%s)", string(resp))
	}
	return words, nil
}
