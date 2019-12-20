package main

import (
	"fmt"
)

func main() {
	// 重点 讲解 / 、 %
	// 说明： 如果运算的数都是整数  那么除后 去掉小数部分 ， 保留整数
	fmt.Println(10 / 4)

	var n1 float32 = 10 / 4 //
	fmt.Println(n1)
	// 希望保留小数部分， 则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)
}
