package main

import "fmt"

type Person struct {
	Num string
}

// 给a 类型绑定一分方法
func (p Person) test() {
	fmt.Println("test() name=", p.Num)
}

func main() {
	var p Person
	p.Num = "tom"
	p.test() // 调用方法
}
