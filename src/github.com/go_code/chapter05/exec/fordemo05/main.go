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
}
