package main

import "fmt"

func testBol() {
	var a bool
	fmt.Println(a)

	a = true
	fmt.Println(a)

	a = !a
	fmt.Println(a)

	var b bool = true
	if a == true && b == true { //与
		fmt.Println("right")
	} else {
		fmt.Println("not right")
	}

	if a == true || b == true { // 或
		fmt.Println("|| right")
	} else {
		fmt.Println("||  not right")
	}
	fmt.Printf("%t %t\n", a, b) // 格式化输出

	//1. 布尔类型
	//a. var b bool 和 var b bool = true 和 var b = false
	//b. 操作符==和 != c. 取反操作符: !b
	//d. &&和||操作符
	//e. 格式化输出占位符: %t
}

func testInt() {
	var a int8 //取值范围在正128到-127之间
	a = 8
	fmt.Println("a=", a)
	a = -12
	fmt.Println("a=", a)

	var b int
	b = 138338338
	fmt.Println("b=", b)
	b = int(a)
	fmt.Println("b=", b)

	//var c float32 = 5.6
	var c = 5.6
	fmt.Println(c)

	fmt.Printf("a=%d a=%x c=%f\n", a, a, c)
	//2. 整数和浮点数类型
	//a. int8、int16、int32、int64
	//b. uint8、uint16、uint32、uint64
	//c. int 和 uint，和操作系统平台相关
	//d. float32 和 float64浮点类型
	//e. 所有整数 初始化为0，所有浮点数初始化为0.0，布尔类型初始化为false

	//3. 整数和浮点数类型
	//a. Go是强类型语⾔言，不不同类型相加以及赋值是不不允许的 b. 那怎么样才能实现，不不同类型相加呢?
	//c. 输出占位符:整数%d，%x⼗十六进制，%f浮点数
}
func main() {
	// testBol()
	testInt()

}
