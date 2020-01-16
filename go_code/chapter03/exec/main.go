package main

import (
	_ "fmt" // 如果我们没有使用到一个包，但是有想去掉， 前面加一个 _  表示忽略
)

// func main() {
	// 练习
	// 	var n1 int32 = 12
	// 	var n2 int64
	// 	var n2 int8
	// 	n2 = int64(n1) + 20 // int32 ---> int64 错误
	// 	n3 = int8(n1) + 20  // int32 --->  int8 错误
	// 	fmt.Println("n2=", n2, "n3=", n3)

	// var n1 int32 = 12
	// var n3 int8
	// var n4 int8
	// n4 = int8(n1) + 127 // [编译通过，但是 结果 不是127+12 按照溢出处理]
	// n3 = int8(n1) + 128 // [编译不通过]
	// fmt.Println(n3)
// }
