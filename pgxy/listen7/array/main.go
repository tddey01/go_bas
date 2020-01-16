package main

import "fmt"

func testArrya1() {
	var a [5]int
	a[0] = 200
	a[1] = 300
	a[2] = 400
	a[3] = 500
	a[4] = 600
	fmt.Println(a)
	//数组初始化
}
func testArrya2() {
	var a [5]int = [5]int{10, 20, 30}
	fmt.Println(a)
	//定义时数组初始化
}

func testArrya3() {
	a := [5]int{10, 20, 30}
	fmt.Println(a)
	//定义时数组初始化
}

func testArrya4() {
	a := [...]int{10, 20, 30}
	fmt.Println(a)
	//定义时数组初始化
}

func testArrya5() {
	a := [3]int{10, 3, 5}
	fmt.Println(a)
	//定义时数组初始化
}

func testArrya6() {
	a := [5]int{3: 100, 4: 200}
	fmt.Println(a)
	//定义时数组初始化
}

func testArrya7() {
	a := [5]int{3: 100, 4: 500}
	fmt.Println(a)
	var b [5]int
	b = a
	fmt.Println(b)
	// 数组⻓长度是类型的⼀一部分
}

func testArrya8() {
	a := [5]int{3: 400, 4: 600}
	fmt.Printf("len(a) = %d\n", len(a))
	// len内置函数
}

func testArrya9() {
	a := [5]int{3: 400, 2: 600}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	//. 数组遍历
}

func testArrya10() {
	a := [5]int{3: 200, 4: 500}
	//var index, value int
	// for index,value := range a{
	//for index, value = range a {
	//	fmt.Printf("a[%d] = %d\n", index, value)
	//}
	for _, value := range a {
		fmt.Printf("%d\n", value)
	}
	//a、b是不不同类型的数组，不不能赋值
}

func testArrya11() {
	var a [3][2]int
	a[0][0] = 10
	a[0][1] = 20
	a[1][0] = 30
	a[1][1] = 40
	a[2][0] = 50
	a[2][1] = 60

	fmt.Println(a)

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
	fmt.Println("other method")

	for i, val := range a {
		fmt.Printf("row[%d]=%v\n", i, val)
		for j, val2 := range val {
			fmt.Printf("(%d,%d) = %d", i, j, val2)
		}
		fmt.Println()
	}
	// ⼆二维数组 两种遍历方式
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func testArrya12() {
	a := [5]int{3: 500, 2: 100}
	b := a
	b[0] = 1000
	fmt.Printf("a=%v\n", a)
	fmt.Printf("b=%v\n", b)
}

func testArray13() {
	var a int = 1000
	b := a
	b = 3000
	fmt.Printf("a=%d b=%d\n", a, b)
}

func modify(b [3]int) {
	b[0] = 1000
}

func testArray14() {
	var a [3]int = [3]int{10, 20, 30}
	modify(a)
	fmt.Println(a)

}
func main() {
	//testArrya1() //数组初始化
	// testArrya2() //定义时数组初始化
	// testArrya3() //定义时数组初始化
	// testArrya4() //定义时数组初始化
	//testArrya5() //定义时数组初始化
	//testArrya6()
	// testArrya7() //数组⻓长度是类型的⼀一部分
	// testArrya8() // len内置函数
	//testArrya9() //. 数组遍历
	//testArrya10() //a、b是不不同类型的数组，不不能赋值
	//testArrya11() // ⼆二维数组

	//a := [3][2] string{
	//	{"lion","tiger"},
	//	{"cat","dog"},
	//	{"pigeon","peacok"},
	//}
	//printarray(a)
	//var b [3][2] string
	//b[0][0] = "apple"
	//b[0][1] = "samsung"
	//b[1][0] = "microsoft"
	//b[1][1] = "google"
	//b[2][0] = "AT&T"
	//b[2][1] = "T-Moblie"
	//fmt.Printf("\n")
	//printarray(b)

	testArrya12()
	testArray13()
	testArray14()

}
