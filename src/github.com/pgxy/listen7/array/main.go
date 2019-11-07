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
func main() {
	//testArrya1() //数组初始化
	// testArrya2() //定义时数组初始化
	// testArrya3() //定义时数组初始化
	// testArrya4() //定义时数组初始化
	//testArrya5() //定义时数组初始化
	//testArrya6()
	// testArrya7() //数组⻓长度是类型的⼀一部分
	// testArrya8() // len内置函数
	testArrya9() //. 数组遍历
}
