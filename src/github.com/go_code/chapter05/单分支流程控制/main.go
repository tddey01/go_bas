package main

import (
	"fmt"
)

func main(){
	// ifDemo.go
	// 编写一个 可以输入年龄 如果该同志年龄大于18 输出 你年龄大于18， 要对自己负责

	/*
	 分析
	 1 年龄 ==> var age  int
	 2 从控制台接收一个输入 fmt.Scanlne(&age)
	 3 if 判断
	*/
	
	var  age  int 
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你年龄大于18， 要对自己负责")
	}
}