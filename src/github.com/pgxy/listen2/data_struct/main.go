package main

import "fmt"

func main() {
	var a bool
	fmt.Println(a)

	a = true
	fmt.Println(a)

	a = !a
	fmt.Println(a)

	var b bool = true
	if a == true && b == true { //与
		fmt.Println("right")
	} else {
		fmt.Println("not right")
	}

	if a == true || b == true { // 或
		fmt.Println("|| right")
	} else {
		fmt.Println("||  not right")
	}
}

//1. 布尔类型
//a. var b bool 和 var b bool = true 和 var b = false
//b. 操作符==和 != c. 取反操作符: !b
//d. &&和||操作符
//e. 格式化输出占位符: %t
