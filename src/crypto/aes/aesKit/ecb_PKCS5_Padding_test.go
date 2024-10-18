package aesKit

//import (
//	"fmt"
//	"testing"
//)
//
//func TestEncryptEcbPKCS5PaddingToString(t *testing.T) {
//	key := []byte("0123456789abcdef")
//	plainText := "强无敌qwdqw`12345678990-=~!@#$%^&*()_+d测试test"
//
//	cipherText, err := EncryptEcbPKCS5PaddingToString(key, []byte(plainText))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("cipherText:", cipherText)
//	plainData1, err := DecryptEcbPKCS5PaddingFromString(key, cipherText)
//	if err != nil {
//		panic(err)
//	}
//	plainText1 := string(plainData1)
//	fmt.Println("plainText1:", plainText1)
//
//	if plainText != plainText1 {
//		panic("not equal")
//	}
//}
