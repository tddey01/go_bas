package main

import "fmt"

func modify(a *int) {
	*a = 10000
}
func main() {
	var a int = 100
	fmt.Printf("before modeify :%d addr:%p\n", a, &a)
	modify(&a)
	fmt.Printf("ofter modeify :%d addr:%p\n", a, &a)
}
