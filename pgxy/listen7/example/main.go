package main

import (
	"fmt"
	"math/rand"
)

func sumArray(a [10]int) int {
	var sum int = 0
	//sum := 0
	//第一种遍历方式
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}

	// 第二种 遍历方式
	//for _,val := range a{
	//	sum = sum + val
	//}
	return sum
}

func testArraySum() {
	var b [10]int
	for i := 0; i < len(b); i++ {
		//b[i] = i
		//产生一个0到999的随机数
		b[i] = rand.Intn(10000)
		//产生一个0到Int最大值的随机数
		//b[i] = rand.Int()
	}
	sum := sumArray(b)
	fmt.Println(sum)
	fmt.Printf("sum=%d\n", sum)
}

func TwoSum(a [5]int, target int) {
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		for j := 0; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d %d)\n", i, j)
			}
		}
	}
}
func testTwoSum() {
	//var b [5]int = [5]int{1,3,5,8,7}
	//b := [5]int{1,3,5,8,7}
	//b := [...]int{1, 3, 5, 8, 7}
	b := [...]int{1, 3, 5, 6, 7}
	TwoSum(b, 8)
}

func main() {
	//testArraySum()
	testTwoSum()
}
