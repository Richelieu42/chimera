package aesKit

import (
	"fmt"
	"testing"
)

func TestEncryptCfbZeroPaddingToString(t *testing.T) {
	plainText := "强无敌qwdqw`12345678990-=~!@#$%^&*()_+d测试test"
	// 16位
	key := []byte("0123456789abcdef")
	iv := []byte("I Love Go Frame!")
	var padding int

	cipherText, err := EncryptCfbZeroPaddingToString([]byte(plainText), key, &padding, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println("cipherText:\n", cipherText)

	plainData1, err := DecryptCfbZeroPaddingFromString(cipherText, key, padding, iv)
	if err != nil {
		panic(err)
	}
	plainText1 := string(plainData1)
	fmt.Println("plainText1:\n", plainText1)

	if plainText != plainText1 {
		panic("not equal")
	}
}
