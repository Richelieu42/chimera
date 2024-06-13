package httpKit

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"net/http"
	"time"
)

// GetCertificateInfo
/*
获取https过期时间
	https://www.topgoer.cn/docs/gochajian/gofdgjh
*/
func GetCertificateInfo(url string) (info *x509.Certificate, err error) {
	if !strKit.StartWith(url, "https://") {
		err = errorKit.Newf("invalid url(%s)", url)
		return
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	info = resp.TLS.PeerCertificates[0]
	return
	//fmt.Println("过期时间:", certInfo.NotAfter)
	//fmt.Println("组织信息:", certInfo.Subject)
}
