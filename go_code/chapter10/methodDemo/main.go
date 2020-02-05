package main

import "fmt"

type Person struct {
	Name string
}

// 1、给Person结构体添加speak 方法 输出 xxx是一个好人
func (p Person) speak() {
	fmt.Println(p.Name, "是一个goodman")
}

// 2、给Person结构体添加jisuan 方法,可以计算从 1+..+1000的结果,
//说明方法体内可以函数一样，进行各种运算
func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结过是=", res)
}

//给Person结构体jisuan2 方法,该方法可以接收一个参数n，计算从 1+..+n 的结果
func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果2=", res)
}

//给Person结构体添加getSum方法,可以计算两个数的和，并返回结果
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2

}

// 给person 类型绑定一分方法
func (person Person) test() {
	person.Name = "jack"
	fmt.Println("test() name=", person.Name)
}

func main() {
	var p Person
	p.Name = "tom"
	p.test() // 调用方法
	fmt.Println("main p.Name=", p.Name)

	p.speak() //好人

	p.jisuan() //计算结果
	p.jisuan2(10)

	res := p.getSum(10, 20)
	fmt.Println("res=", res)
}
