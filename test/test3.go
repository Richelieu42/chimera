package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/token/jwtKit"
)

func main() {
	key := []byte("c27022ed-f258-46ca-b894-6d54e5db4ce7")

	m, err := jwtKit.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ3Mjg2MzIsImlhdCI6MTcxNDEyMzgzMiwiY2hhbm5lbCI6ImV4YW1wbGUifQ.i0QbtLifViBGTjnN6h78Ts3CE7z0hm62bsWaWRsTmpA", func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return key, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}
