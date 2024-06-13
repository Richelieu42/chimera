package payKit

import "github.com/go-pay/gopay"

func GetVersion() string {
	return gopay.Version
}
