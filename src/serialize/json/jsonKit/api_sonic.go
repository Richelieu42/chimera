//go:build (linux || windows || darwin) && sonic && avx && go1.20 && (amd64 || arm64)

package jsonKit

import (
	"github.com/bytedance/sonic"
	"github.com/klauspost/cpuid/v2"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v3/src/core/osKit"
	"github.com/sirupsen/logrus"
)

/*
!!!: 并非 amd64 CPU 就行了，还需要支持 avx指令集 等.（e.g.yozo某台amd64内网机就不行）
*/
func init() {
	library = "bytedance/sonic"
	defaultApi = sonic.ConfigDefault
	stdApi = sonic.ConfigStd

	if !cpuKit.HasFeature(cpuid.AVX) {
		logrus.WithFields(logrus.Fields{
			"os":   osKit.OS,
			"arch": osKit.ARCH,
		}).Fatalf("[%s, JSON, SONIC] AVX isn't supported by CPU!!!", consts.UpperProjectName)
		return
	}

	testAPI()
}
