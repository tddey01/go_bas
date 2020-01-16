package main

import (
	"fmt"
	"unsafe"
)

// Golang 中bool类型使用
func main() {
	var b = false
	fmt.Println("b=", b)
	// 注意事项
	// 1 ， bool；类型占用存储空间是1个字节
	fmt.Println("b 占用空间 = ", unsafe.Sizeof(b))
	// 2 . bool类型只能取true 或者false

}
