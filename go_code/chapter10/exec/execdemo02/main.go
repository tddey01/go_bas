package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type Stu Student

type integer int

func main() {
	var stu1 Student
	var stu2 Stu
	// stu2 = stu1  正确 答案 错误  可以这样强转 	stu2 = Stu(stu1)
	stu2 = Stu(stu1)
	fmt.Println(stu1, stu2)

	var i integer = 10
	var j int = 20
	// j =i  争取吗 错误  可以这样强转  j = int(i)
	j = int(i)
	fmt.Println(j)

}
