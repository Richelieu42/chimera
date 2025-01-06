//go:build go1.20 && arm64 && sonic

package jsonKit

import (
	"github.com/bytedance/sonic"
)

func init() {
	library = "bytedance/sonic"
	defaultApi = sonic.ConfigDefault
	stdApi = sonic.ConfigStd
}
