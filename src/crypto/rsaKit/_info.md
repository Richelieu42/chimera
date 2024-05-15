## 简介
RSA加密算法是一种广泛使用的 "非对称加密" 方法.

## PKCS1 VS PKCS8
早期openssl1.0之前版本，一般提供是pkcs#1格式，有很多软件只支持pkcs#1格式（js rsa模块)，那么你可以选择生成该种类型。
现在一般流行是pkcs#8格式.

## 生成公钥私钥
#### 在线生成
在线生成非对称加密公钥私钥对
    http://web.chacuo.net/netrsakeypair
#### 通过openssl命令生成
# 生成私钥
openssl genrsa -out private.pem 2048
# 生成公钥
openssl rsa -in private.pem -outform PEM -pubout -out public.pem

## 参考
Golang-RSA加密解密-数据无大小限制
    https://www.cnblogs.com/akidongzi/p/12036165.html
golang里面private key证书的加解密
    https://www.jianshu.com/p/c102a639cc50
Golang 常见加密算法实现
    https://mp.weixin.qq.com/s/dg-9Ew1zy64Pgvok6zA-2w
