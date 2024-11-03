package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// 创建一个文件服务器，将其根目录设置为"./static"
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fileServer)
	// 定义一个API端点的处理函数
	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from API!"))
	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
