package main

import (
	"fmt"
)

func main() {
	// 输出10句 ”你好 大笨猪“
	// fmt.Println("你好 大笨猪")
	// fmt.Println("你好 大笨猪")
	// fmt.Println("你好 大笨猪")
	// fmt.Println("你好 大笨猪")
	// fmt.Println("你好 大笨猪")
	// fmt.Println("你好 大笨猪")
	// fmt.Println("你好 大笨猪")
	// fmt.Println("你好 大笨猪")
	// fmt.Println("你好 大笨猪")
	// fmt.Println("你好 大笨猪")
	// golang中 有循环控制语句来处理循环的执行某段代码的方式--》for循环
	// for循环快速入门

	for i := 1; i <= 10; i++ {
		// fmt.Println("你好 大笨猪", i)
		fmt.Printf("%v 你好 大笨猪\n", i)
	}

	/*
		对for循环来说 有四要素
			循环变量初始化
			循环条件
			循环操作（语句） 有人也叫循环体
			循环变量迭代

		for循环执行的顺序说明
			执行循环变量初始
			执行循环条件判断
			如果循环条件为真，就执行循环操作
			执行循环变量迭代
			反复执行2，3，4，步骤，指导循环条件为False 就推出for循环
	*/

	// for 循环的第二种写法
	j := 1 // 循环变量初始化
	for j <= 10 {
		fmt.Printf("%v 你好 大笨猪\n", j)
		j++ //循环变量迭代
	}
}
