package main

import (
	"strings"
)

var morseCodeMap = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".", 'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-", 'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--", 'Z': "--..",
	'1': ".----", '2': "..---", '3': "...--", '4': "....-", '5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.", '0': "-----",
	' ': "/",
}

var reverseMorseCodeMap = map[string]rune{}

func init() {
	for char, code := range morseCodeMap {
		reverseMorseCodeMap[code] = char
	}
}

func encrypt(text string) string {
	var result strings.Builder
	for _, char := range text {
		if code, ok := morseCodeMap[char]; ok {
			result.WriteString(code)
			result.WriteString(" ")
		}
	}
	return result.String()
}

func decrypt(morseText string) string {
	words := strings.Fields(morseText)
	var result strings.Builder
	for _, word := range words {
		if char, ok := reverseMorseCodeMap[word]; ok {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func main() {
	originalText := "你好，世界！"
	encrypted := encrypt(originalText)
	decrypted := decrypt(encrypted)
	println("原始文本：", originalText)
	println("加密后的摩尔斯密码：", encrypted)
	println("解密后的文本：", decrypted)
}
