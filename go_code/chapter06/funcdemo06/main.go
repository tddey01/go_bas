package main

import "fmt"

var (
	// fun1 就是一个全局匿名函数
	// 第四种方法使用

	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 在定义匿名函数时 就直接调用，这种方式匿名函数只能使用一次
	//第一种使用方法
	// 案例 求两个数的和 使用匿名函数的方式完成
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	// 第二种使用方法
	// 将你匿名函数  func(n1 int, n2 int) int赋给 a变量
	// 则a 的数据类型就是函数类型 此时 我们可以通过该a完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(11, 13)
	fmt.Println("res2=", res2)
	res3 := a(90, 13)
	fmt.Println("res2=", res3)

	// 全局匿名函数的调用
	res4 := Fun1(20, 80)
	fmt.Println("res4=", res4)

}
