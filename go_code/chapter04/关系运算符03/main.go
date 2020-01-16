package main

import (
	"fmt"
)

func main() {
	//关系运算符使用
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2) // 不等于 false
	fmt.Println(n1 != n2) // 大于 true
	fmt.Println(n1 > n2)  // 小于 true
	fmt.Println(n1 >= n2) // 小于等于 true
	fmt.Println(n1 < n2)  // 大于 false
	fmt.Println(n1 <= n2) // 大于等于 false
	flag := n1 > n2
	fmt.Printf("flang=%v\n", flag)
	// 关系运算符的结果都是bool类型，也就是要么是true  要么是false
	// 关系运算符组成的表达式 我们成为 关系表达式 a > b
	// 比较运算符  ”==“ 不能误写成 ”=“
}
