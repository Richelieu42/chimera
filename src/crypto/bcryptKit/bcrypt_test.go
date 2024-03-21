package bcryptKit

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	plainPwd := "qwdqwdqwdq强无敌wd"

	hashedPwd, err := HashPassword([]byte(plainPwd))
	if err != nil {
		panic(err)
	}
	fmt.Println("hashedPwd:", string(hashedPwd)) // hashedPwd: $2a$10$wqkTtz2ZWBYhZUo6knRNf.bYRTA/Mjv27XMuwNJJB2AKZZocAyIxi

	fmt.Println(ComparePasswords(hashedPwd, []byte(plainPwd)) == nil) // true
}
