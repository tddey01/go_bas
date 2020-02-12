package main

import (
	"fmt"
)

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i=", i)
}
func main() {
	var stu Stu // 结构体变量 实现了 Say() 实现了 AInterface
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say() // integer Say i = 10
}
