package httpKit

import (
	"fmt"
	"testing"
)

func TestGetCertificateInfo(t *testing.T) {
	url := "https://www.baidu.com"

	info, err := GetCertificateInfo(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("NotBefore:", info.NotBefore)
	fmt.Println("NotAfter:", info.NotAfter)
	fmt.Println("Issuer:", info.Issuer)
	fmt.Println("Subject:", info.Subject)
}
