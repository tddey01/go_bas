package main

import (
	"fmt"
	"strconv"
	"time"
)

func test3() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
func main() {
	// 执行test3前， 先获取当亲unix时间戳
	start := time.Now().Unix()
	test3()
	end := time.Now().Unix()
	fmt.Printf("执行test3()耗费时间为%v\n",end-start)
}
