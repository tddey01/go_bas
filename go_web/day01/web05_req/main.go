package main

import (
	"fmt"
	"net/http"
)

// 创建处理函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求地址是:", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求地址后的查询字符串是:", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值", r.Header.Get("Accept-Encoding"))
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
