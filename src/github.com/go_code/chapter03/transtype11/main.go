package main

import (
	"fmt"
)

// Golang 中 基本数据类型的转换

func main() {

	var i int32 = 100
	// 希望将 i => float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v n1=%v n2=%v n3=%v\n", i, n1, n2, n3)

}
