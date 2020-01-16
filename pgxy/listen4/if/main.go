package main

import "fmt"

func testif1() {
	num := 10
	if num%2 == 0 {
		fmt.Printf("num :%d is even\n", num)
	} else {
		fmt.Printf("num :%d  is add\n", num)
	}
	// if else语句句块
}

func testif2() {
	num := 10
	if num > 5 && num < 10 {
		fmt.Printf("num :%d is > $ and < 10\n", num)
	} else if num >= 10 && num < 20 {
		fmt.Printf("num:%d is > 10 and < 20\n", num)
	} else if num >= 20 && num < 30 {
		fmt.Printf("num :%d is > 20 and < 30\n", num)
	} else {
		fmt.Printf("num :%d is > 20 and >30\n", num)
	}
	// if else语句句
}

func testif3() {
	if num := 11; num%2 == 0 {
		fmt.Printf("num :%d is even\n", num)
	} else {
		fmt.Printf("num :%d  is add\n", num)
	}
}

func getNum() int {
	return 10
}

func testIf4() {
	//num := 10
	//var num int
	//num = 10
	if num := getNum(); num%2 == 0 {
		fmt.Printf("num:%d is even\n", num)
	} else {
		fmt.Printf("num:%d is odd\n", num)
	}

	//fmt.Printf("num=%d\n", num)
}

func main() {
	//testif1()
	// testif2()
	// testif3()
	testIf4()
}
