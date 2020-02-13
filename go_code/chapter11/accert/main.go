package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point //ok
	// 如何将 a 付给一个Point变量
	var b Point
	//b = a // 可以吗 --》 error 错误提示
	b = a.(Point)  // 可以正常使用  类型断言
	fmt.Println(b)
}
