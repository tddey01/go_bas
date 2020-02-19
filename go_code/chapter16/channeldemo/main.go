package main

import "fmt"

func main() {
	// 案例管道使用方法
	//1、 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2 看看intChan是什么？
	fmt.Printf("intChan 的值=%v intChan本身的地址%p\n", intChan, &intChan)

	// 3 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 20

	//intChan <- 21  // 注意事项  当我们给管道写入数据时，不能超过其容量

	//4 看看管道的长度和cap（容量）
	fmt.Printf("channel len=%v  cap=%v \n", len(intChan), cap(intChan)) // 2 3

	//5 从管道里面读取数据
	var num2 int
	num2 = <-intChan
	fmt.Printf("num2=%v\n", num2)
	fmt.Printf("channel len=%v  cap=%v \n", len(intChan), cap(intChan))

	// 6 在没有携程的情况下， 如果我们的管道数据已经全部取出，在取就会报错 deadlock
	num3 := <-intChan
	num4 := <-intChan
	fmt.Printf("num3=%v   num4=%v\n", num3, num4)

	num5 := <-intChan
	fmt.Printf("num5=%v\n", num5)

}
