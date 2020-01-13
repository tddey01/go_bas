package main

import "fmt"

// 函数中的变量是局部的 函数外不生效
func test() {
	//n1 是 test函数的局部变量  只能在test函数里面使用
	// var n1 int = 10

}

func test02(n1 int) {
	n1 = n1 + 10
	fmt.Println("test02()=", n1)
}
func main() {
	n1 := 20
	test02(n1)
	fmt.Println("main() n1=", n1)
}
