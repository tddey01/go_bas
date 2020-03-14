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

// 包含
func testTemptale(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t := template.Must(template.ParseFiles("template1.html", "template2.html"))
	// 执行
	t.Execute(w, "我能在两个文件中显示吗")
}

func testDefine(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t := template.Must(template.ParseFiles("define.html"))
	// 执行
	t.ExecuteTemplate(w, "model", "")
}

// 测试 testDefine2
func testDefine2(w http.ResponseWriter, r *http.Request) {
	age := 17
	var t *template.Template
	if age < 18 {
		// 解析模板
		// t = template.Must(template.ParseFiles("define2.html", "content2.html"))
		t = template.Must(template.ParseFiles("define2.html"))
	} else {
		// 解析模板
		t = template.Must(template.ParseFiles("define2.html", "content1.html"))
	}
	// 执行
	t.ExecuteTemplate(w, "model", "")
}

func main() {
	// 条件
	http.HandleFunc("/testIf", testIf)
	// 迭代
	http.HandleFunc("/testRange", testRange)
	// 设置
	http.HandleFunc("/testwith", testwith)
	// 包含
	http.HandleFunc("/testTemplate", testTemptale)
	// 模板
	http.HandleFunc("/testdefine", testDefine)

	// 模板
	http.HandleFunc("/testdefine2", testDefine2)

	http.ListenAndServe(":8080", nil)
}
