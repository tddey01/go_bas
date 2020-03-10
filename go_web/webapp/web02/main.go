package main

import (
	"fmt"
	"net/http"
)

// 结构体
type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "我们自己创建的处理器请求！")
}
func main() {
	//
	myHandler := MyHandler{}
	// 处理请求
	http.Handle("/myHandler", &myHandler)
	// 路由分发
	http.ListenAndServe(":8080", nil)
}
