package main

import (
	"fmt"
)

func main() {
	// const  a int  = 100  	 // 常量定义 必须赋值
	// const b string = "hello"

	const (
		a  int = 100
		b string = "hello"    // 优雅写法
	)

	// const (  // 更加专业写法
	// 	a = iota
	// 	b
	// 	c
	// )
	// const (
	// 	a = 1 << iota
	// 	b 
	// 	c
	// )
	fmt.Printf("a = %d b = %s \n",a,b)
}
