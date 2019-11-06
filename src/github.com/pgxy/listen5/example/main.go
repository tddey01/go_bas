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

func is_shuixianhua(n int) bool {
	first := n % 10
	second := (n / 100) % 10
	third := (n / 100) % 10

	fmt.Printf(" n%d first:%d second:%d third:%d\n", n, first, second, third)
	return false
}

func example2() {
	for i := 100; i < 1000; i++ {
		if is_shuixianhua(i) == true {
			fmt.Printf("[%d] is shuixianhua", i)
		}

	}
}
func main() {
	//example1()
	example2()
}
