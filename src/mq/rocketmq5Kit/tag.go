package rocketmq5Kit

import "github.com/richelieu42/chimera/v2/src/core/sliceKit"

func MixTags(tags ...string) string {
	return sliceKit.Join(tags, "||")
}

func GetTagString(tag *string) string {
	if tag == nil {
		return ""
	}
	return *tag
}
