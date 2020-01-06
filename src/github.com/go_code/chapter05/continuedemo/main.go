package main

import "fmt"

func main() {
	// continue案例
	// 演示执行标签的形式来使用continue

	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			continue
	// 			fmt.Println("ok")
	// 		}
	// 		fmt.Println("j=", j)
	// 	}
	// }

	// for i := 0; i < 2; i++ {
	// 	for j := 1; j < 4; j++ {
	// 		if j == 2 {
	// 			continue
	// 		}
	// 		fmt.Println("i=", i, "j=", j)
	// 	}
	// }
here:
	for i := 0; i < 2; i++ {
		for j := 1; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
}
