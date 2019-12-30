package main

import (
	"fmt"
)

func main() {
	//分析思路
	// score 分数变量 int
	// 选择多分支流程控制
	// 键盘输入成绩 fmt.Scanln

	var score int
	fmt.Println("请输入成绩:")
	fmt.Scanln(&score)

	// 多分支判断
	if score == 100 {
		fmt.Println("奖励一辆宝马一台")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台iPhone7")
	} else if score >= 60 && score <= 80 {
		fmt.Println("奖励一个ipad")
	} else {
		fmt.Println("什么都不奖励")
	}

	// 使用陷阱
	var n int = 10
	if n > 9 {
		fmt.Println("ok1")
	} else if n > 6 {
		fmt.Println("ok2")
	} else if n > 3 {
		fmt.Println("ok3")
	} else {
		fmt.Println("ok")
	}
}
