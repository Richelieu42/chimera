package jwtKit

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
)

func TestJwt(t *testing.T) {
	secret := []byte("c27022ed-f258-46ca-b894-6d54e5db4ce7")
	claims := jwt.MapClaims{
		"c": "y",
	}

	jwtStr, err := Sign(secret, jwt.SigningMethodHS256, claims)
	if err != nil {
		panic(err)
	}
	fmt.Println("jwt:", jwtStr)

	m, err := Verify(jwtStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("m:", m)
}

func TestJwtWithCaesar(t *testing.T) {
	secret := []byte("c27022ed-f258-46ca-b894-6d54e5db4ce7")
	caesarShift := 10

	claims := jwt.MapClaims{
		"c": "y",
	}

	jwtStr, err := SignWithCaesar(caesarShift, secret, jwt.SigningMethodHS256, claims)
	if err != nil {
		panic(err)
	}
	fmt.Println("jwt:", jwtStr)

	m, err := VerifyWithCaesar(caesarShift, jwtStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("m:", m)
}
