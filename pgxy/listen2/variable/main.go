package main

import (
	"fmt"
)
 
func main() {
	/*
	var a int
	var b bool
	var c string
	var d float32
	*/

	var (
		a int
		b bool
		c string
		d float32
	)
	fmt.Printf("a = %d  b = %t  c=%s  d=%f \n ", a, b, c, d)
	/*  %d 输出一个整数   %t 输出布尔类型 %s 字符串类型   %f 是浮点数  \n是换行符 */
	a = 10
	b = true
	c = "hello"
	d = 10.0

	fmt.Printf("a = %d  b = %t  c=%s  d=%f ", a, b, c, d)
}
