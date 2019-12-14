package main

import (
	"fmt"
)

// Golang 中字符类型使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0' // 字符的o

	// 当我嗯直接输出byte值， 就是输出了对应的字符的码值
	// 'a' ==>
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	// 如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c  c2=%c\n", c1, c2)

	// var c3 byte = '北' // overflows byte 溢出
	var c3 int = '北' 
	fmt.Printf("c3=%c c3对应的码值=%d\n", c3,c3)

}
