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
}
