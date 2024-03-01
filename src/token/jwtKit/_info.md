## 参考
- 


## golang-jwt/jwt
- [github 6.1k Star](https://github.com/golang-jwt/jwt)

#### 使用的是 base64.RawURLEncoding，生成的 jwt 可以直接放到 url 中
详见源码: Token.SignedString()

#### 默认情况下（不加密JWT），前两部分可以通过 base64.RawURLEncoding 解码
```go
package main

import (
	"encoding/base64"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/crypto/base64Kit"
)

func main() {
	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.-BRTwjN-sAlUjO-82qDrNHdMtGAwgWH05PrN49Ep_sU"
	fmt.Println(str)

	// {"alg":"HS256","typ":"JWT"} <nil>
	fmt.Println(base64Kit.DecodeStringToString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", base64Kit.WithEncoding(base64.RawURLEncoding)))
	// {"foo":"bar","nbf":1444478400} <nil>
	fmt.Println(base64Kit.DecodeStringToString("eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9", base64Kit.WithEncoding(base64.RawURLEncoding)))
	// �S�3~�  T��ڠ�4wL�`0�a������)�� <nil>
	fmt.Println(base64Kit.DecodeStringToString("-BRTwjN-sAlUjO-82qDrNHdMtGAwgWH05PrN49Ep_sU", base64Kit.WithEncoding(base64.RawURLEncoding)))
}
```
