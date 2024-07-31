//go:build go1.20 && arm64

package jsonKit

import (
	"github.com/bytedance/sonic"
)

func init() {
	library = "bytedance/sonic"
	defaultApi = sonic.ConfigDefault
	stdApi = sonic.ConfigStd

	//if !cpuKit.HasFeature(cpuid.AVX) {
	//	text := fmt.Sprintf("AVX isn't supported with os(%s) and arch(%s)", osKit.OS, osKit.ARCH)
	//	panic(text)
	//	return
	//}

	testAPI()
}
