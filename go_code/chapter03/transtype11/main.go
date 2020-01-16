package main

import (
	"fmt"
)

// Golang 中 基本数据类型的转换

func main() {

	var i int32 = 100
	// 希望将 i => float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v n1=%v n2=%v n3=%v\n", i, n1, n2, n3)

	// 被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化
	fmt.Printf("i type is %T\n", i) // int32

	//在转换中，比如将 int64 转成 int8 【-128---127】 ，编译时不会报错，只是转换的结果是按 溢出处理，和我们希望的结果不一样。 因此在转换时，需要考虑范围.
	var num1 int64 = 999999
	var num2 int32 = int32(num1) // 转换的结果是按 溢出处理
	fmt.Println("num2=", num2)
}
