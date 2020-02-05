package main

import "fmt"

type Person struct {
	Name string
}

// 给a 类型绑定一分方法
func (p Person) test() {
	p.Name = "jack"
	fmt.Println("test() name=", p.Name)
}

func main() {
	var p Person
	p.Name = "tom"
	p.test() // 调用方法
	fmt.Println("main p.Name=", p.Name)
}
