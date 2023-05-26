package rsaKit

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
)

func parsePKCS1PrivateKey(s []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(s)
	if block == nil {
		return nil, errorKit.Simple("private key error(%s)", "PKCS1")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func parsePKCS8PrivateKey(s []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(s)
	if block == nil {
		return nil, errorKit.Simple("private key error(%s)", "PKCS8")
	}

	keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return keyInterface.(*rsa.PrivateKey), nil
}
