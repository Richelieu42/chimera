package ocrKit

import (
	"fmt"
	httpClientKit2 "github.com/richelieu-yang/chimera/v3/src/component/web/httpClient/httpClientKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"github.com/richelieu-yang/chimera/v3/src/urlKit"
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
	base64Str, err := base64Kit.EncodeFileToString(imagePath)
	if err != nil {
		return nil, err
	}
	params := map[string][]string{
		"language_type": {"CHN_ENG"},
		"image":         {urlKit.EncodeURIComponent(base64Str)},
	}

	// 发请求
	_, respData, err := httpClientKit2.Post(url, httpClientKit2.WithPostParams(params))
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	if err := jsonKit.Unmarshal(respData, &m); err != nil {
		return nil, err
	}

	// 解析响应
	words, err := parseMapToWords(m)
	if err != nil {
		return nil, err
	}
	if words == nil {
		return nil, errorKit.Newf("failure response(%s)", string(respData))
	}
	return words, nil
}
