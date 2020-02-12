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

type Binterface interface {
	Hello()
}

type Monster struct{

}
func (m Monster) Hello(){
	fmt.Println("Monster Hello~~")
}

func (m Monster) Say(){
	fmt.Println("Monster Say~~")
}
func main() {
	var stu Stu // 结构体变量 实现了 Say() 实现了 AInterface
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say() // integer Say i = 10

	//Monter 实现了AInterface和BInterface
	var monster Monster
	var a2 AInterface = monster
	var b2 Binterface =  monster
	a2.Say()
	b2.Hello()

}
