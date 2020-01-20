package main

import (
	"fmt"
	"time"
)

func test() {
	// 使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数， 可以捕获到异常
		if err != nil {  // 说捕获异常错误
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}
func main() {
	// 测试
	test()
	for {
		fmt.Println("main() 下吗代码...")
		time.Sleep(time.Second)
	}

}
