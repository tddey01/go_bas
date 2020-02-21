package main

import (
	"fmt"
	"reflect"
)

//通过反射，修改,
// num int 的值
// 修改 student的值

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	// 看看 rVal的Kind是
	fmt.Printf("rVal Kind=%v\n", rVal.Kind())
	//2 rVal
	rVal.Elem().SetInt(20)
}

func main() {

	var num int = 10
	reflect01(&num)

	fmt.Println(num)

	// // 可以这样理解rVal.Elem()
	// num := 9
	// ptr *int =&num
	// num2 := *ptr // === 类似于 rVal.Elem()
}
