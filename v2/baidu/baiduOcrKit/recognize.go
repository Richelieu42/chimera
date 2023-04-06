package baiduOcrKit

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/core/errorKit"
	"github.com/richelieu42/chimera/v2/core/file/fileKit"
	"github.com/richelieu42/chimera/v2/jsonKit"
	"github.com/richelieu42/chimera/v2/urlKit"
	"github.com/richelieu42/chimera/v2/web/httpClientKit"
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
	url := fmt.Sprintf("%s?access_token=%s", "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic", token.AccessToken)

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
	resp, err := httpClientKit.SimplePost(url, params)
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
