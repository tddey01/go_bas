package main

import (
	"fmt"
)

type Animal interface {
	Talk()
	Eat()
	Name() string
}
type Dog struct {
}

func (d Dog) Talk() {
	fmt.Println("汪汪")
}
func (d Dog) Eat() {
	fmt.Println("我在吃东西")
}

func (d Dog) Name() string {
	fmt.Println("我的名字")
	return "旺财"
}

type Pig struct {
}

func (d Pig) Talk() {
	fmt.Println("坑坑")
}
func (d Pig) Eat() {
	fmt.Println("我在吃草")
}

func (d Pig) Name() string {
	fmt.Println("我的名字")
	return "猪八戒"
}

func main() {
	var d Dog
	var a Animal
	a = d

	a.Eat()
	a.Talk()
	a.Name()

	var p Pig
	a = p
	a.Eat()
	a.Talk()
	a.Name()

}
