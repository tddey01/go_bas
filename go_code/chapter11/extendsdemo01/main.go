package main

import "fmt"

// 编写一个学生考试系统
type Student struct {
	Name  string
	Age   int
	Score int
}

// 将Pupil 和Graduate 共有的方法也绑定到 *Student
//显示他的成绩
func (stu *Student) ShowInfo() {
	fmt.Printf("学生的名字=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	//业务判断
	stu.Score = score
}

// 增加一个做算术题方法
func (stu *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

//小学生
type Pupil struct {
	Student // 嵌入了Student匿名的结构体
}

// 这是 Pupil结构体特有的方法保留
func (p *Pupil) testing() {
	fmt.Println("小学生考试中....")
}

// 大学生考试 研究生考试..
type Graduate struct {
	Student // 嵌入了Student匿名的结构体
}

func (p *Graduate) testing() {
	fmt.Println("大学生考试中....")
}
func main() {
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()
	fmt.Println("res=", pupil.Student.GetSum(10, 24))

	graduate := &Graduate{}
	graduate.Student.Name = "jack"
	graduate.Student.Age = 28
	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.ShowInfo()
	fmt.Println("res2=", graduate.Student.GetSum(30, 50))
}
