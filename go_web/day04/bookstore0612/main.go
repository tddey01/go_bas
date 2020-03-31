package main

import (
	"go_bas/go_web/day04/bookstore0612/controller"
	"html/template"
	"net/http"
)

// 处理器函数
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//	执行
	t.Execute(w, "")
}

func main() {
	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/main", IndexHandler)

	// 去登陆
	http.HandleFunc("/login", controller.Login)
	// 注册
	http.HandleFunc("/regist", controller.Regist)
	// 通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)

	// 获取所有图书
	http.HandleFunc("/getBooks", controller.GetBooks)
	// 添加图书
	//http.HandleFunc("/addBook", controller.AddBook)
	////删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	// 去更新图书页面
	http.HandleFunc("/toUpdateBookPage",controller.ToUpdateBookPage)
	//// 更新图书
	//http.HandleFunc("/updateBook",controller.UpdateBook)
	// 更新图书 或者 添加图书
	http.HandleFunc("/updateOraddBook",controller.UpdateOrAddBook)

	http.ListenAndServe(":8080", nil)
}
