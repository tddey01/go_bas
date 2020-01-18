package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("num1的类型%T num1的值=%v num1的内存地址%v\n", num1, num1, &num1)

	num2 := new(int) // *int 指针类型
	// num2的类型%T  --> *int
	// num2的值=%v   -->  根据当前系统使用内存状态来分配内存地址
	// num2的内存地址%v  -->
	// 指向的值=%v  --> 0
	*num2 = 100
	fmt.Printf("num2的类型%T num2的值=%v num2的内存地址%v 指向的值=%v\n", num2, num2, &num2, *num2)
}
