package main

import "fmt"

func testFor2() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i-%d\n", i)
	}
}
func main() {
	testFor2()
}
