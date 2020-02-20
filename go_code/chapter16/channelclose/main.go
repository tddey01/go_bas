package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) // close 关闭 panic: send on closed channel
	// 这是不能够在写入数据到channel管道里面
	// intChan<-400
	fmt.Println("ok")
	// 当这个管道关闭后 可以正常读取管道里面的数据
	n1 := <-intChan
	fmt.Println(n1)

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 // 放入100个数据到这个管道里面
	}
	// 遍历 01
	// for i := 0; i < len(intChan2); i++ {
	// 	fmt.Println(i)
	// }  //禁止这种遍历
	//在遍历时，如果channel没有关闭，则会出现deadlock的错误
	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
