package pinyinKit

import (
	"github.com/mozillazg/go-pinyin"
)

func PinYin(str string) [][]string {
	args := pinyin.NewArgs()

	return pinyin.Pinyin(str, args)
}

func PinYinWithTone(str string) [][]string {
	args := pinyin.NewArgs()
	args.Style = pinyin.Tone

	return pinyin.Pinyin(str, args)
}

func PinYinWithToneAndHeteronym(str string) [][]string {
	args := pinyin.NewArgs()
	args.Style = pinyin.Tone
	// 开启多音字模式
	args.Heteronym = true

	return pinyin.Pinyin(str, args)
}
