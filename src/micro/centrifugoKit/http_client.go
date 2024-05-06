package centrifugoKit

import (
	"crypto/tls"
	"net/http"
	"time"
)

var DefaultHttpClient = &http.Client{
	Timeout: time.Second * 3,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}
