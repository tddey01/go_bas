package main

import "fmt"

// 定义全局变量
var n4 = 100
var n5 = 200
var name1 = "jack"

// 上面的声明方式 也可以改成一次性声明
var (
	m1    = 300
	m2    = 900
	name2 = "mary"
)

func main() {
	//该案案例演示golang如何一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 一次性声明多个变量方式2
	var s1, name, s2 = 100, "kn", 888
	fmt.Println("s1=", s1, "name=", name, "s2=", s2)

	// 一次性声明多个变量方式3 同样可以使用类型推导
	a1, a2, a3 := 100, "kn", 999
	fmt.Println("a1=", a1, "a2=", a2, "a3=", a3)

	//输出全局变量
	fmt.Println("n4=", n4, "n5=", n5, "name=", name1)

	fmt.Println("m1=", m1, "m2=", m2, "name2=", name2)
}
