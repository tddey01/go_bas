package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "hello World!", r.URL.Path)
	fmt.Fprintln(w, "通过自己创建多路复用器请求", r.URL.Path)
}

func main() {
	//创建多路复用器
	mux := http.NewServeMux()
	// 处理请求
	//http.HandleFunc("/", handler)
	mux.HandleFunc("/", handler)
	//创建路由
	//http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", mux)
}
