package main

import (
	"fmt"
	"go_bas/go_code/chapter10/factory/model"
)

func main() {
	// 创建要给Student实例
	var stu = model.Student{
		Name:  "tom",
		Score: 78.9,
	}
	fmt.Println(stu)

	// /定student结构体是首字母小写，我们可以通过工厂模式来解决
	var stu1 = model.NewStudent("tom", 88.8)
	fmt.Println(*stu1) //&{....}

	fmt.Println("name=", stu1.Name, " score=", stu1.GetScore())

}
