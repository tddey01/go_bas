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

	if b == 100 {
		var c int = 100
		fmt.Printf("c=%d\n", c)
	}
	if d := 100; d > 0 {
		fmt.Printf("d=%d", d)
	} else {
		fmt.Printf("else d=%d", d)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("for i=%d\n", i)
	}

}

func main() {
	// testGlobalVariable()
	testlocalVariable()
}
