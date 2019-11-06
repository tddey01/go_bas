package main

import "fmt"

var a int = 100

// 全局变量量，在程序整个⽣生命周期有效
func testGlobalVariable() {
	fmt.Printf("a=%d", a)
}

func testlocalVariable() {
	var b int = 100
	fmt.Printf("a=%d b=%d\n", a, b)
}

func main() {
	// testGlobalVariable()
	testlocalVariable()
}
