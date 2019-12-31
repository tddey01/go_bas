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

	// 案例四
	// 大家都知道，男大当婚，女大当嫁。那么女方家长要嫁女儿，当然要提出一定的条件：
	//高：180cm以上；富：财富1千万以上；帅：是。条件从控制台输入。
	// 如果这三个条件同时满足，则：“我一定要嫁给他!!!”
	// 如果三个条件有为真的情况，则：“嫁吧，比上不足，比下有余。”
	// 如果三个条件都不满足，则：“不嫁！”

	// var height int32  | var money float32 | var handsome bool

	//分析思路
	//1. 应该设计三个变量 var height int32  | var money float32 | var handsome bool
	//2. 而且需要从终端输入 fmt.Scanln
	//3. 使用多分支if--else if -- else

	var height int32
	var money float32
	var handsome bool

	fmt.Println("请输入身高(厘米)")
	fmt.Scanln(&height)
	fmt.Println("请输入财富(千万)")
	fmt.Scanln(&money)
	fmt.Println("请输入是否帅(true/false)")
	fmt.Scanln(&handsome)

	if height > 180 && money > 1.0 && handsome {
		fmt.Println("我一定要嫁给他!!!")
	} else if height > 180 || money > 1.0 && handsome {
		fmt.Println("嫁吧，比上不足，比下有余")
	} else {
		fmt.Println("不嫁")
	}
}

