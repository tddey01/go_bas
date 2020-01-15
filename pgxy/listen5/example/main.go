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

	//fmt.Printf(" n%d first:%d second:%d third:%d\n", n, first, second, third)
	sum := first*first*first + second*second*second + third*third*third
	if sum == n {
		return true
	}
	return false
}

func example2() {
	for i := 100; i < 1000; i++ {
		if is_shuixianhua(i) == true {
			fmt.Printf("[%d] is shuixianhua", i)
		}

	}
}

func calc(str string) (charCount int, numCount int, spaceCount int, otherCount int) {
	uifCchars := []rune(str)

	for i := 0; i < len(uifCchars); i++ {
		if uifCchars[i] >= 'a' && uifCchars[i] <= 'z' || uifCchars[i] >= 'A' && uifCchars[i] <= 'Z' {
			charCount++
			continue
		}

		if uifCchars[i] >= '0' && uifCchars[i] < '9' {
			numCount++
			continue
		}

		if uifCchars[i] == ' ' {
			spaceCount++
			continue
		}

		otherCount++
	}
	return
}

func example3() {
	var str string = "adhhk      我爱天安门1276845"
	charCount, numCount, spConut, other := calc(str)
	fmt.Printf("char conut :%d num conut:%d sp conut %d other Conut:%d", charCount, numCount, spConut, other)
}
func main() {
	//example1()
	//example2()
	example3()
}
