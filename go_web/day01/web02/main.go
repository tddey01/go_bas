package main

import (
	"fmt"
	"net/http"
	"time"
)

// 结构体
type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "我们自己创建的处理器请求！")
	fmt.Fprintln(w, "通过详细配置服务器的信息处理器请求！")
}
func main() {
	//
	myHandler := MyHandler{}
	// 处理请求
	// http.Handle("/myHandler", &myHandler)
	// 创建server结构 详细配置里面字段
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		ReadTimeout: time.Second * 2,
	}
	// 路由分发
	//http.ListenAndServe(":8080", nil)
	server.ListenAndServe()
}
