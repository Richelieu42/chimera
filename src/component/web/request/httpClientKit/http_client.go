package httpClientKit

import (
	"crypto/tls"
	"net/http"
	"time"
)

// DefaultHttpClient 默认的 *http.Client 实例
var DefaultHttpClient *http.Client = NewHttpClient(time.Second*3, true)

func NewHttpClient(timeout time.Duration, insecureSkipVerify bool) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecureSkipVerify,
			},
		},
	}
}
