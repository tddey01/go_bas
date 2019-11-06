package main

import "fmt"

func testswitch1() {
	var a int = 2
	if a == 1 {
		fmt.Printf("a=1\n")
	} else if a == 2 {
		fmt.Printf("a=2\n")
	} else if a == 3 {
		fmt.Printf("a=3\n")
	} else if a == 4 {
		fmt.Printf("a=4\n")
	} else {
		fmt.Printf("a=5\n")
	}
}

func getValue() int {
	return 1
}
func testswitch2() {
	// var a  int = 7
	// switch a {
	switch a := getValue(); a {
	case 1:
		fmt.Printf("a=1\n")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	case 4:
		fmt.Printf("a=4\n")
	case 5:
		fmt.Printf("a=5\n")
	case 6:
		fmt.Printf("a=6\n")
	default:
		fmt.Printf("invalid a = %d\n", a)
	}
	// 如果想让某个分支执行完后，继续执行其后面的分支，可以用 fallthrough 语句。
}
func testswitch3() {
	switch a := getValue(); a {
	case 1, 2, 3, 4, 5:
		fmt.Printf("a>=1 and <=5\n")
	case 6, 7, 8, 9:
		fmt.Printf("a >= 6 and a <= 9\n")
	default:
		fmt.Printf("invalid a = %d\n", a)
	}
}
func main() {
	// testswitch1()
	//testswitch2()
	testswitch3()
}
