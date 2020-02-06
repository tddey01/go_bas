package main

import "fmt"

/*
学生案例：
编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float64类型。
结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值。
在main方法中，创建Student结构体实例(变量)，并访问say方法，并将调用结果打印输出。

*/
type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int
	Score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 Name=[%v] Gender=[%v] Age=[%v] Id=[%v] Score=[%v]",
		student.Name, student.Gender, student.Age, student.Id, student.Score)
	return infoStr
}
func main() {
	//测试
	// 创建一个student实例变量
	var stu = Student{
		Name:   "tom",
		Gender: "male",
		Age:    22,
		Id:     1000,
		Score:  99.98,
	}
	str := stu.say()
	fmt.Println(str)
}
