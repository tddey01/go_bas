package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	for num := 1; num <= 100000; num++ {
		flag := true

		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// primeChan <- num
		}
	}
	end := time.Now().Unix()
	fmt.Println("使用普通的方法耗时=", end-start)
}
