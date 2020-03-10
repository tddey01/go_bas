package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello World!", r.URL.Path)
}

func main() {
	// 处理请求
	http.HandleFunc("/", handler)
	//创建路由
	http.ListenAndServe(":8080", nil)
}
