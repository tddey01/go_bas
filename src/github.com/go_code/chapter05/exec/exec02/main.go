package main

import (
	"fmt"
	"math"
)

func main() {
	//分析思路
	// score 分数变量 int
	// 选择多分支流程控制
	// 键盘输入成绩 fmt.Scanln

	// var score int
	// fmt.Println("请输入成绩:")
	// fmt.Scanln(&score)

	// // 多分支判断
	// if score == 100 {
	// 	fmt.Println("奖励一辆宝马一台")
	// } else if score > 80 && score <= 99 {
	// 	fmt.Println("奖励一台iPhone7")
	// } else if score >= 60 && score <= 80 {
	// 	fmt.Println("奖励一个ipad")
	// } else {
	// 	fmt.Println("什么都不奖励")
	// }

	// // 使用陷阱
	// var n int = 10
	// if n > 9 {
	// 	fmt.Println("ok1")
	// } else if n > 6 {
	// 	fmt.Println("ok2")
	// } else if n > 3 {
	// 	fmt.Println("ok3")
	// } else {
	// 	fmt.Println("ok")
	// }

	// var b bool = true
	// if b == false { // 如果写成 b = false 能编译通过吗？ 如果能， 结果是？ 编译错误
	// 	fmt.Println("a")
	// } else if b {
	// 	fmt.Println("b") //b 结果是b
	// } else if !b {
	// 	fmt.Println("c") //c
	// } else {
	// 	fmt.Println("d")
	// }

	// 求ax2+bx+c=0方程的根 a、b、c分别为函数的额参数  如果 b2-4ac>0 则有两个解
	// b2-4ac=0 则有一个解 b2-4ac<0 则2无解
	// 提示1  x1=(-b+sqrt(b2-4ac))/2a
	//		 x2=(-b-sqrt(b2-4ac))/2a
	//  提示math.Sqrt(num); 可以求平方根 需要引入math包

	// 分析思路
	/*
		a b c 是三个float64
		使用到给出的是数学公式
		使用多分支
		使用math。Squr方法 => 手册
	*/

	var a float64 = 2.0
	var b float64 = 4.0
	var c float64 = 2.0

	m := b*b - 4*a*c

	//多分支判断
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v x2=%v\n", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v\n", x1)

	} else {
		fmt.Println("无解...")
	}
}
