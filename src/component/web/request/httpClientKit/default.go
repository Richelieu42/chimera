package httpClientKit

import (
	"crypto/tls"
	"net/http"
	"time"
)

// DefaultHttpClient 默认的 *http.Client 实例
var DefaultHttpClient *http.Client = &http.Client{
	Timeout: time.Second * 3,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}
