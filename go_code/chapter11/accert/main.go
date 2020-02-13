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
	b = a.(Point) // 可以正常使用  类型断言
	fmt.Println(b)

	// 类型断言的其他案例

	var x interface{}
	var b2 float32 = 1.1
	x = b2 // 空接口 可以接收任意类型
	// x => float32 【使用类型断言】
	//y , ok := x.(float64)
	//y , ok := x.(float32)
	//y := x.(float32)  // 只能断言 初始化类型
	//if ok == true{
	if y, ok := x.(float32); ok { // 高级写法
		fmt.Println("convert success")
		fmt.Printf("yT的类型是 %T 值是=%v\n", y, y)
	} else {
		fmt.Println("convert fail")
	}
	//fmt.Printf("yT的类型是 %T 值是=%v\n",y,y)
	fmt.Println("继续执行...")
}
