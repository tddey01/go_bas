package main

import "fmt"

func justify(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func example1() {
	for i := 2; i < 100; i++ {
		if justify(i) == true {
			fmt.Printf("[%d] is prime", i)
		}

	}
}
func main() {
	example1()

}
