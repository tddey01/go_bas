package main

import (
	"fmt"
)

// Golang 中标识符的使用

func main() {

	// Golang 中严格区分大小写
	// Golang 中 认为 num 和 Num 是不同的变量
	var num int = 10
	var Num int = 20
	fmt.Printf("num=%v Num=%v\n", num, Num)

	// 标识符不能包含空格
	// var ab c int =  30 //syntax error 语法错误

	// _ 是空标志符，用于占用
	// var  _ int = 40 //error
	// fmt.Println(_) // cannot use _ as value 不可使用

	// 语法可以通过， 但是要求不能使用
	var int int = 90
	fmt.Println(int)
}

//标识符的案例
// hello // ok
// hello12 //ok
// 1hello // error ,不能以数字开头 h-b // error ,不能使用 -
// x h // error, 不能含有空格
// h_4 // ok
// _ab // ok
// int // ok , 我们要求大家不要这样使用
// float32 // ok , 我们要求大家不要这样使用
// _   // error
// Abc // ok
