package main

import (
	"fmt"
	"time"
)

//intChan放入 1- 8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 800000; i++ {
		intChan <- i
	}
	// 关闭intChan
	close(intChan)
}

//intChan取出数据， 并判断是否为素数，如果是，就放入primeChan里面
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	// 使用for循环
	// var num int
	var flag bool
	for {
		// time.Sleep(time.Millisecond * 100)
		num, ok := <-intChan
		flag = true
		if !ok { // intChan 取不到。。
			break
		}
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			// 将这个数就放入到primeChan
			primeChan <- num
		}
	}
	fmt.Println("有一个primeChan 携程因为取不到数据 退出")
	//这里我们还不能关闭 primeChan
	//向 exitChan 写入一个标志 true
	exitChan <- true
}
func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000) // 放入结果
	// 表示 退出的管理
	exitChan := make(chan bool, 4) // 4 个

	start := time.Now().Unix()
	// 开启一个携程 intChan放入 1- 8000个数
	go putNum(intChan)
	// 开启四个携程， 从intChan取出数据， 并判断是否为素数，如果是，就放入primeChan里面

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里我们主线程，进行处理
	//直接
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)
		// 当我们exitChan 取出了4个结果， 就可以放心的关闭了 primeChan
		close(primeChan)
	}()

	// 遍历我们的 primeNum 结果 ， 把结果取出来
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		// 将结果输出
		// fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("主线程退出")
}
