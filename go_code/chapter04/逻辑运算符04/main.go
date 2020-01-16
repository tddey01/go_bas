package main

import "fmt"

// 声明一个函数（测试）
func test() bool {
	fmt.Println("test....")
	return true
}

func main() {

	var i int = 10
	// 短路与
	// 说明 因为 i<9 为false  因此后面的test() 就不执行了
	if i < 9 && test() {
		fmt.Println("ok...")
	}

	// 说明 因为 i<9 为true  因此后面的test() 就不执行了
	if i < 9 || test() {
		fmt.Println("hello...")
	}
	/*
		// 逻辑运算符的使用   &&  与
		var age int = 40
		if age > 30 && age < 50 {
			fmt.Println("ok1")
		}

		if age > 30 && age < 40 {
			fmt.Println("ok2")
		}

		// 逻辑运算符的使用   ||  或
		if age > 30 || age < 50 {
			fmt.Println("ok3")
		}

		if age > 30 || age < 40 {
			fmt.Println("ok4")
		}

		// 逻辑运算符的使用 ！
		if age > 30 {
			fmt.Println("ok5")
		}
		if !(age > 30) {
			fmt.Println("ok6")
		}
	*/
}
