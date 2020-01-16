package main

import "fmt"

func testScanf() {
	var a int
	var b string
	var c float32
	var str string = "88 hello  8.8"
	// fmt.Scanf("%d %s %f", &a, &b, &c) //监听键盘输入
	fmt.Sscanf(str, "%d%s%f\n", &a, &b, &c)
	fmt.Printf("a:%d b:%s c:%f\n", a, b, c)
}

func testScan() {
	var a int
	var b string
	var c float32

	fmt.Scan(&a, &b, &c) // 换行符 作为分隔符
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func testScanln() {
	var a int
	var b string
	var c float32
	var str string = "88\n\n hello\n  8.8"
	fmt.Sscanln(str, &a, &b, &c) //遇到换行符结束输入
	fmt.Printf("a=%d b=%s c=%f\n", a, b, c)
}

func main() {
	// testScanf()
	// testScan()
	testScanln()
}
