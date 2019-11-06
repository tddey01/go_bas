package main

import "fmt"

func sayhello() {
	fmt.Printf("hellow word\n")
}

func add(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	//sayhello()
	s := add(100, 300)
	fmt.Println(s)
}
