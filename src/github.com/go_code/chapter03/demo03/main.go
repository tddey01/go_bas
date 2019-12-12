package main

import "fmt"

func main() {
	//该案案例演示golang如何一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// 一次性声明多个变量方式2
	var s1, name, s2 = 100, "kn", 888
	fmt.Println("s1=", s1, "name=", name, "s2=", s2)

	// 一次性声明多个变量方式3 同样可以使用类型推导
	a1, a2, a3 := 100, "kn",999
	fmt.Println("a1=",a1, "a2=",a2, "a3=",a3)
}
