package main

import (
	"fmt"
	"github.com/pgxy/listen11/calc"
)

func main() {
	// var sum = calc.Add(2, 5)
	sum = calc.Add(2, 5)
	fmt.Printf("sum = %d\n", sum)
}

//go run运⾏行行go代码， 如果有多个⽂文件，需要把所有⽂文件都写到go run后⾯面
//go build 编译go代码，如果是可执⾏行行程序，默认会在当前⽬目录⽣生成可执⾏行行 程序，可以使⽤用-o指定可执⾏行行程序⽣生成的⽬目录
//go install编译go代码，并且把可执⾏行行程序拷⻉贝到GOPATH的bin⽬目录，⾃自定义或 第三⽅方包会拷⻉贝到GOPATH的pkg⽬目录
