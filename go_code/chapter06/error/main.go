package main

import (
	"errors"
	"fmt"
)

func test() {
	// 使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数， 可以捕获到异常
		if err != nil {  // 说捕获异常错误
			fmt.Println("err=", err)
			// 这里就可以将错误信息发给管理员
			fmt.Println("发送邮件给admin@sohu.com")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func test1() {
	// 使用defer + recover 来捕获和处理异常
	defer func() {
		if err := recover(); err != nil { // 说捕获异常错误	// recover()内置函数， 可以捕获到异常
			fmt.Println("err=", err)
			// 这里就可以将错误信息发给管理员
			fmt.Println("发送邮件给admin@sohu.com")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

// 函数去读取已配置文件init.conf的信息
// 如果文件名传入不正确，我们就返回一个自顶一个的错误信息
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取
		return nil
	} else {
		//返回一个自定错误
		return errors.New("读取文件错误")
	}
}
func test02() {
	err := readConf("config2.ini")
	if err != nil {
		//如果读取文件发送错误， 就输出这个错误， 并终止程序
		panic(err)
	}
	fmt.Println("test02 继续执行...")
}
func main() {
	// 测试
	// test()
	// test1()
	// for i := 0; i <= 3; i++ {
	// 	fmt.Println("main() 下吗代码...")
	// 	time.Sleep(time.Second)
	// }
	//测试自定错误的使用
	test02()
	fmt.Println("main()下吗的代码..")
}
