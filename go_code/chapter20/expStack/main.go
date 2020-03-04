package main

import (
	"errors"
	"fmt"
)

// 使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int     // 表示我们栈最最大可以存放的个数
	Top    int     // 表示栈顶, 因为栈顶固定，因此我们直接使用Top
	arr    [20]int // 数组模拟栈

}

//入栈
func (this *Stack) Push(val int) (err error) {
	// 先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	// 放入
	this.arr[this.Top] = val
	return
}

//遍历栈，注意需要从栈顶开始遍历
func (this *Stack) List() {
	// 先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	// 不为空 从栈顶开始输出
	fmt.Println("栈的情况如下：")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	//	判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	//先取值，再 this.Top--
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

// 判断一个字符是不是一个运算符[+, -, *, /]
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 运算的方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

//编写一个方法，返回某个运算符的优先级[程序员定义]
//[* / => 1 + - => 0]
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res

}
func main() {
	// 数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//	符号栈
	openStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	exp := "3+2*6-2"
	//	定义一个index， 帮助扫描exp
	index := 0

	//for 循环
	for {
		ch := exp[index : index+1] // 字符串.

	}
}
