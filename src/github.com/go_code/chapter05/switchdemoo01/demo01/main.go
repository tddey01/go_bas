package main

import (
	"fmt"
)

func tests(char byte) byte {
	return char + 1
}
func main() {
	// 	案例：
	// 请编写一个程序，该程序可以接收一个字符，比如: a,b,c,d,e,f,g
	// a表示星期一，b表示星期二 …  根据用户的输入显示相依的信息.

	// 要求使用 switch 语句完成
	// 1 定义一个变量接收字符
	// 2 使用switch 完成

	// var key byte
	// fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	// fmt.Scanf("%c", &key)

	// switch tests(key) + 1 { //讲语法现象
	// case 'a':
	// 	fmt.Println("周一 猴子穿新衣")
	// case 'b':
	// 	fmt.Println("周二 猴子当小二")
	// case 'c':
	// 	fmt.Println("周三 猴子爬雪山")
	// case 'd':
	// 	fmt.Println("周四")
	// case 'e':
	// 	fmt.Println("周五")
	// case 'f':
	// 	fmt.Println("周六")
	// case 'g':
	// 	fmt.Println("周日")
	// default:
	// 	fmt.Println("输入有误...")
	// }

	/*
		switch 细节说明
		  1、case 后是一个表达式(即：常量值、变量、一个有返回的函数等都可以)
		  2、case 后的各个表达式的值数据类型，必须和switch的表达式数据类型一直
		  3、case 后面可以带多个表达式，及用逗号间隔。比如case表达式1，表达式2...
		  4、case 后面的表达式如果是常量值(字面量)，则需要求不能重复
		  5、case h后面不需要呆break，程序匹配到一个case后就会执行对应的代码块，然后退出switch，如果一个匹配不到，则执行default
		  6、default 语句不是必须的
	*/
	//

	// var n1 int32 = 20
	// var n2 int64 = 20

	// switch n1 {
	// case n2: //错误 原因是 n2的数据类型和n1 不一致
	// 	fmt.Println("ok1")
	// default:
	// 	fmt.Println("没有匹配到")
	// }

	// var n1 int32 = 5
	// var n2 int32 = 20

	// switch n1 {
	// case n2, 10, 5: //case  后面可以有多个表达式
	// 	fmt.Println("ok1")
	// default:
	// 	fmt.Println("没有匹配到")
	// }

	// switch 后面也可以表达式 类似 if --else 分支来使用
	// var age int = 25
	// switch {
	// case age == 10:
	// 	fmt.Println("age == 10")
	// case age == 20:
	// 	fmt.Println("age == 20")
	// default:
	// 	fmt.Println("没有匹配到")
	// }

	// case 中也可以对  范围进行判断
	var score int = 80
	switch {
	case score > 90:
		fmt.Println("成绩优秀..")
	case score >= 70 && score <= 90:
		fmt.Println("成绩优良...")
	case score >= 60 && score < 70:
		fmt.Println("成绩及格...")
	default:
		fmt.Println("不及格...")
	}

	// switch 后也可以直接声明/定义一个变量，分好结束，不推荐使用

	switch grade := 95; { // 在golang 中  可以这样写
	case grade > 90:
		fmt.Println("成绩优秀..~")
	case grade >= 70 && grade <= 90:
		fmt.Println("成绩优良...~")
	case grade >= 60 && grade < 70:
		fmt.Println("成绩及格...~")
	default:
		fmt.Println("不及格...~")
	}

	// switch  的穿透  fallthrought
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默认只能穿透一次
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}
}
