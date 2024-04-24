package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/token/jwtKit"
	"time"
)

func main() {
	//str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.GvmVuP_7yADlqHk6fB7Tcq2V5EGY98PQw3EkX3DbBmQ"
	secret := "16bfd798-4f9f-4362-98e8-d88cb4997db2"

	j := jwtKit.NewJWT([]byte(secret))
	method := jwt.SigningMethodHS256
	//method := jwt.SigningMethodHS384
	//method := jwt.SigningMethodHS512

	tokenString, err := j.Sign(method, map[string]interface{}{
		"a": "b",
		//"exp": jwt.NewNumericDate(time.Now().Add(-time.Hour)),
		"iat": jwt.NewNumericDate(time.Now().Add(-time.Hour)),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("tokenString:", tokenString)

	//mc, err := j.Verify(str, func(token *jwt.Token) (interface{}, error) {
	//	// Don't forget to validate the alg is what you expect:
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	//	}
	//
	//	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	//	return j.Key, nil
	//})
	//if err != nil {
	//	if jwtKit.IsTokenExpiredError(err) {
	//		panic("Token is expired.")
	//	}
	//	panic(err)
	//}
	//fmt.Println(mc)
	//fmt.Println(mc.GetIssuedAt())
}
