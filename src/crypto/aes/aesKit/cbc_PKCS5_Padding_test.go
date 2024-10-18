package aesKit

import (
	"fmt"
	"testing"
)

func TestEncryptCbcPKCS5PaddingToString(t *testing.T) {
	plainText := "强无敌qwdqw`12345678990-=~!@#$%^&*()_+d测试test"
	// 16位
	key := []byte("0123456789abcdef")
	iv := []byte("I Love Go Frame!")

	cipherText, err := EncryptCbcPKCS5PaddingToString([]byte(plainText), key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Println("cipherText:\n", cipherText)

	plainData1, err := DecryptCbcPKCS5PaddingFromString(cipherText, key, iv)
	if err != nil {
		panic(err)
	}
	plainText1 := string(plainData1)
	fmt.Println("plainText1:\n", plainText1)

	if plainText != plainText1 {
		panic("not equal")
	}
}
