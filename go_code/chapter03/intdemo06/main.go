package main

import (
	"fmt"
	"unsafe"
)

// golang中整数类型使用
func main() {
	var i int = 1
	fmt.Println("i=", i)

	// 测试一下int8的范围 -128 ~ 127
	// 其他的 int16,int32,int64以此类推...
	var j int8 = 127
	fmt.Println("j=", j)

	// 测试一下  uint8的范围 ，其它的 nint16， uint32， uint64 一样
	var k uint8 = 255
	fmt.Println("k=", k)

	// int  , uint , rune, byte的使用
	var a int = 8900
	fmt.Println("a=", a)
	var b uint = 1
	fmt.Println("b=", b)
	var c byte = 255
	fmt.Println("c=", c)

	//整数的使用细节
	var n1 = 100 //? n1 是什么类型
	//如何查看某个变量数据类型
	// fmt.Printf()， 可以用于做格式化输出
	fmt.Printf("n1 的 类型 %T\n", n1)

	// 如何在程序查看某个变量的占用字节大小和数据类型(使用多)
	var n2 int64 = 10
	//unsafe.Sizeof(n2) 是unsafe包的一个函数， 可以返回n2变量占用字节数
	fmt.Printf("n2 de 类型 %T n2 占用的字节数是 %d\n", n2, unsafe.Sizeof(n2))

	// Golang 程序中整数类型变量在使用时，遵守保小不保大的原则
	// 即：在保证程序正确云心下， 尽量是占用空间小的数据类型
	var age byte = 4
}
