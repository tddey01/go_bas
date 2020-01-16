package main

import (
	"fmt"
)

func main() {

	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr *int
	ptr = &num
	*ptr = 10 // 这里修改时 会到num的值变化
	fmt.Printf("num=%v\n", num)

	var a int = 300
	var b int = 400
	var prt *int = &a
	*prt = 100
	prt = &b
	*prt = 200
	fmt.Printf("a=%d b=%d *prt=%d\n", a, b, *prt)
	// a=100 b=200 *prt=200
}
