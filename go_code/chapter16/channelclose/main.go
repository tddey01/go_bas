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
	n1 := <- intChan
	fmt.Println(n1)
}
