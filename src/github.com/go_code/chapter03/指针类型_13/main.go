package main

import (
	"fmt"
)

// Golang  中指针类型
func main() {
	// 基本数据类型在内存布局
	// 基本数据类型，变量存的就是值，也叫值类型
	var i int = 10
	// i 地址是多少？
	// 获取变量的地址，用&，比如: var num int, 获取 num 的地址:&num
	fmt.Println("i的地址=", &i)

	// 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值
	// 比如:var ptr *int = &num
	// 1、 ptr 是一个指针变量
	// 2、 ptr 的类型 *int
	// 3、 ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr =%v\n", ptr)
	fmt.Printf("ptr 的地址=%v\n", &ptr)
	fmt.Printf("ptr 指针的值=%v\n", *ptr)
}
