package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/*
		我们为了生成一个随机数，还需要设置rand设置一个种子
	*/
	// rand.Seed(time.Now().Unix())
	// //如何随机生成1-100整数
	// n := rand.Intn(100) + 1 // 0 100
	// fmt.Println(n)

	//随机生成1-100的一个数，直到生成了99这个数，看看你一共用了几次
	//分析思路：
	//编写一个无限循环的控制，然后不停的随机生成数，当生成了99时，就退出这个无限循环==》break
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99 一共使用了", count)
}
