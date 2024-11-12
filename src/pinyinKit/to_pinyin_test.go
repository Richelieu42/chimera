package pinyinKit

import (
	"fmt"
	"testing"
)

func TestPinYin(t *testing.T) {
	fmt.Println(PinYin("好奇")) // [[hao] [qi]]
}

func TestPinYinWithTone(t *testing.T) {
	fmt.Println(PinYinWithTone("好奇")) // [[hǎo] [qí]]
}

func TestPinYinWithToneAndHeteronym(t *testing.T) {
	fmt.Println(PinYinWithToneAndHeteronym("好奇")) // [[hǎo hào] [qí jī ǎi yǐ]]
}
