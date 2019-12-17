package main

import (
	"fmt"
	"strconv"
)

// Golang 中string转成基本数据类型

func main() {

	var str string = "true"

	var b bool // 来接收 这个变量是空值

	// b, _ = strconv.ParseBool(str)
	// 说明
	// 1 strconv.ParseBool(str) 函数会返回两个值 (value bool, err error)
	// 2 因为我只想获取到value bool ，不想获取err 所以使用 “_” 下划线忽略掉
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "1234590"
	var n1 int64
	var n2 int
	var n3 int16
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	n3 = int16(n1)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	fmt.Printf("n1 type %T n1=%v\n", n2, n2)
	fmt.Printf("n1 type %T n1=%v\n", n3, n3)

	var str3 string = "123.456"
	var f1 float64
	var f2 float32
	f1, _ = strconv.ParseFloat(str3, 64)
	f2 = float32(f1)
	fmt.Printf("n1 type %T n1=%v\n", f1, f1)
	fmt.Printf("n1 type %T n1=%v\n", f2, f2)

	// 注意
	var str4 string = "hello"
	var n4 int64 = 11
	n4, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n4 type %T n1=%v\n", n4, n4)
}
