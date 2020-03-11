package main

import (
	"encoding/json"
	"fmt"
	"go_bas/go_web/day02/web01_req/model"
	"net/http"
)

// 创建处理函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求地址是:", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求地址后的查询字符串是:", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值", r.Header.Get("Accept-Encoding"))
	// 获取请求体内容的长度
	//len := r.ContentLength
	//// 创建byte切片
	//body := make([]byte, len)
	//// 将请求体中的内容康读到boy中
	//r.Body.Read(body)
	//// 在浏览器中显示请求体中的内容
	//fmt.Fprintln(w, "请求体内容有", string(body))

	////解析表单，在调用r.Form之前必须执行该操作
	//r.ParseForm()
	////err := r.ParseForm()
	////fmt.Println("错误信息是", err)
	//// 获取请求参数
	////如果form表单的action属性的URL地址中也有与form表单参数名相同的请求参数，
	////那么参数值都可以得到，并且form表单中的参数值在ULR的参数值的前面
	//fmt.Fprintln(w, "请求参数有", r.Form)
	//fmt.Fprintln(w, "POST请求的form表单中的请求参数有:", r.PostForm)
	//	通过直接调用FormValue方法和PostFormValue方法直接获取请求参数
	fmt.Fprintln(w, "URL中的user请求参数值是：", r.FormValue("user"))
	fmt.Fprintln(w, "Form中的username请求参数值是：", r.PostFormValue("username"))

}

// 相应客户报文定义类型
func testJonsRes(w http.ResponseWriter, r *http.Request) {
	// 设置相应内容类型
	w.Header().Set("Conten-Type", "aapplication/json")
	// 创建User
	user := model.User{
		Id:       1,
		Username: "admin",
		Passord:  "123456",
		Email:    "admin@atguigiu.com",
	}
	// 将User转换json格式
	json, _ := json.Marshal(user)
	// 将json格式返回给客户端
	w.Write(json)
}

// 设置响应 Location属性
func testRedire(w http.ResponseWriter, r *http.Request) {
	// 设置响应头中的Location属性
	w.Header().Set("Location", "https://www.baidu.com")
	w.Header().Set("Server", "google")
	//设置响应的状态码
	w.WriteHeader(302)
}
func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", testJonsRes)
	http.HandleFunc("/testRedirect", testRedire)
	http.ListenAndServe(":8080", nil)
}
