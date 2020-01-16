package main

import (
	"fmt"
)

// Golang 中 string类型使用

func main() {
	// string的基本类型
	var address string = "北京长城 110 hello world"
	fmt.Println(address)

	//字符串一旦赋值了，字符串就不能修改了:在 Go 中字符串是不可变的。
	// var str = "hello"
	// str[0] = 'a' // 这里就不能修改str的内容 即go中的字符串是不可变的

	//字符串的两种表示形式 双引号, 会识别转义字符 反引号
	// 以字符串的原生输出， 包括换行和特殊符号，可以实现防止攻击、
	// 输出源码代码等效果
	str2 := "abc\nabc"
	fmt.Println(str2)

	str3 := `
	func main() {
		// string的基本类型
		var address string = "北京长城 110 hello world"
		fmt.Println(address)
	
		//字符串一旦赋值了，字符串就不能修改了:在 Go 中字符串是不可变的。
		// var str = "hello"
		// str[0] = 'a' // 这里就不能修改str的内容 即go中的字符串是不可变的
		
		//字符串的两种表示形式 双引号, 会识别转义字符 反引号
		// 以字符串的原生输出， 包括换行和特殊符号，可以实现防止攻击、
		// 输出源码代码等效果
		str2 := "abc\nabc"
		fmt.Println(str2)
	}
	`
	fmt.Println(str3)

	// 字符拼接方式
	var str4 = "hello" + "world"
	str4 += "haha!"
	fmt.Println(str4)

	// 当一个拼接的操作很长是， 怎么办， 可以分行写
	var str5 = "hello" + "world" + "hello" + "world" + "hello" +
		"world" + "hello" + "world" + "hello" + "world" + "hello" +
		"world" + "hello" + "world" + "hello" + "world"
	fmt.Println(str5)

	var a int          // 0
	var b float32      //0
	var c float64      //0
	var isMarried bool //false
	var name string    // ""
	// 这里的%v 表示按照变量的值输出
	fmt.Printf("a=%d b=%v c=%v isMarried=%v name=%v \n", a, b, c, isMarried, name)
}
