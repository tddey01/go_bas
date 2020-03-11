package main

import (
	"html/template"
	"net/http"
)

// 创建 处理器函数
func testTemplate(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	//t, _ := template.ParseFiles("index.html")
	// 通过Must函数让Go帮我们自动处理异常
	t := template.Must(template.ParseFiles("index1.html", "index2.html"))
	// 执行
	//t.Execute(w, "hello Template")
	// 将相应数据在index2.html文件显示
	t.ExecuteTemplate(w, "index1.html", "我要去index2.html中")
}

func main() {
	http.HandleFunc("/testTemplate", testTemplate)

	http.ListenAndServe(":8080", nil)
}
