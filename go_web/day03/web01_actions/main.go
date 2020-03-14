package main

import (
	"go_bas/go_web/day03/web01_actions/model"
	"html/template"
	"net/http"
)

// 测试 if
func testIf(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t := template.Must(template.ParseFiles("if.html"))
	age := 17
	// 执行
	t.Execute(w, age > 18)
}

//  测试 range
func testRange(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t := template.Must(template.ParseFiles("range.html"))
	var emps []*model.Employee
	emp := &model.Employee{
		ID:       1,
		LastName: "李晓娜",
		Email:    "lixiaona@jnl.com",
	}
	emps = append(emps, emp)
	emp1 := &model.Employee{
		ID:       2,
		LastName: "李三",
		Email:    "lisan@jnl.com",
	}
	emps = append(emps, emp1)
	emp2 := &model.Employee{
		ID:       3,
		LastName: "李四",
		Email:    "lisi@jnl.com",
	}
	emps = append(emps, emp2)
	//  执行
	t.Execute(w, emps)
}

// 测试 with
func testwith(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t := template.Must(template.ParseFiles("with.html"))
	// 执行
	t.Execute(w, "狸猫")
}

func testTemptale(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t := template.Must(template.ParseFiles("template1.html", "template2.html"))
	// 执行
	t.Execute(w, "我能在两个文件中显示吗")
}
func main() {
	http.HandleFunc("/testIf", testIf)

	http.HandleFunc("/testRange", testRange)
	//
	http.HandleFunc("/testwith", testwith)
	//
	http.HandleFunc("/testTemplate", testTemptale)

	http.ListenAndServe(":8080", nil)
}
