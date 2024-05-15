app=test
path=test/test.go
cgo=0

# 制作者
builder=$(shell whoami)
# 制作日期
build_date=$(shell date '+%FT%T %A %Z')

prepare:
	@go mod tidy
	@go-bindata -fs -o=internal/resources/bindata.go -pkg=resources _resources/...

all: prepare windows-amd64 darwin-arm64 linux-amd64 linux-amd64-noavx linux-arm64 linux-loong64 linux-mips64 linux-mips64le

windows-amd64:
	@rm -rf $(app)-windows-amd64.exe
	@rm -rf $(app)-windows-amd64.upx
	@CGO_ENABLED=$(cgo) GOOS=windows GOARCH=amd64 go build -tags=jsoniter -o $(app)-windows-amd64.exe $(path)
	@upx -9 $(app)-windows-amd64.exe

darwin-arm64:
	@rm -rf $(app)-darwin-arm64
	@CGO_ENABLED=$(cgo) GOOS=darwin GOARCH=arm64 go build -tags=jsoniter -o $(app)-darwin-arm64 $(path)

linux-amd64-noavx:
	@rm -rf $(app)-linux-amd64-noavx
	@rm -rf $(app)-linux-amd64-noavx.upx
	@CGO_ENABLED=$(cgo) GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o $(app)-linux-amd64-noavx $(path)
	@upx -9 $(app)-linux-amd64-noavx

linux-amd64:
	@rm -rf $(app)-linux-amd64
	@rm -rf $(app)-linux-amd64.upx
	@CGO_ENABLED=$(cgo) GOOS=linux GOARCH=amd64 go build -tags="sonic avx" -o $(app)-linux-amd64 $(path)
	@upx -9 $(app)-linux-amd64

linux-arm64:
	@rm -rf $(app)-linux-arm64
	@rm -rf $(app)-linux-arm64.upx
	@CGO_ENABLED=$(cgo) GOOS=linux GOARCH=arm64 go build -tags=jsoniter -o $(app)-linux-arm64 $(path)
	@upx -9 $(app)-linux-arm64

linux-mips64:
	@rm -rf $(app)-linux-mips64
	@CGO_ENABLED=$(cgo) GOOS=linux GOARCH=mips64 go build -tags=jsoniter -o $(app)-linux-mips64 $(path)

linux-mips64le:
	@rm -rf $(app)-linux-mips64le
	@CGO_ENABLED=$(cgo) GOOS=linux GOARCH=mips64le go build -tags=jsoniter -o $(app)-linux-mips64le $(path)

linux-loong64:
	@rm -rf $(app)-linux-loong64
	@CGO_ENABLED=$(cgo) GOOS=linux GOARCH=loong64 go build -tags=jsoniter -o $(app)-linux-loong64 $(path)
