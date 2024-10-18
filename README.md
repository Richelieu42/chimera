# richelieu-yang/chimera

Golang的工具库（Tools of Golang）

[//]: # (![qrcode.png]&#40;qrcode.png&#41;)
<img src="./qrcode.png" alt="qrcode" width="300" height="300">

## 安装此依赖

命令: go get github.com/richelieu42/chimera/v3

## 配置文件（仅供参考）

路径: _chimera-lib/config.yaml

## Linux环境下，自动设置 GOMAXPROCS 的值，以便更好地利用容器的CPU资源

```go
package main

import (
	_ "go.uber.org/automaxprocs"
)

```

