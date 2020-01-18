package main

import "fmt"

func test() {
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}
func main() {
	// 测试
	test()
	fmt.Println("main() 下吗代码...")
}
