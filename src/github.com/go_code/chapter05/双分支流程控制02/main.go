package main

import (
	"fmt"
)

func main() {
	// 编写一个 可以输入年龄 如果该同志年龄大于18 输出 你年龄大于18， 要对自己负责, 否则 ，输出 你的年龄不大这次就放过你
	/*
	 分析
	 1 年龄 ==> var age  int
	 2 从控制台接收一个输入 fmt.Scanlne(&age)
	 3 if  --- else
	*/
	var age int
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你年龄大于18， 要对自己负责")
	} else {
		fmt.Println("你的年龄不大这次就放过你")
	}

}
