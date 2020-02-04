package main

import "fmt"

type A struct {
	Num int
}

type B struct {
	Num int
}

func main() {
	var a A
	var b B
	// a = b    // 不能b的结构体赋值给a
	a = A(b) // 两个结构体 必须有相同的字段 名称 字段类型 字段个数 可以转换
	fmt.Println(a, b)
}
