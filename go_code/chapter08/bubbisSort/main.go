package main

import "fmt"

// 冒泡排序
// func BubblieSort(arr *[5]int) {
// 	fmt.Println("排序钱arr=", (*arr))
// 	temp := 0 // 临时变量（用户做交换）
// 	// 完成第一轮排序(外层1次)
// 	for j := 0; j < 4; j++ {
// 		if (*arr)[j] > (*arr)[j+1] {
// 			// 交换
// 			temp = (*arr)[j]
// 			(*arr)[j] = (*arr)[j+1]
// 			(*arr)[j+1] = temp
// 		}
// 	}
// 	fmt.Println("第一次排序后arr=", (*arr))
// 	// 完成第二轮排序(外层1次)
// 	for j := 0; j < 3; j++ {
// 		if (*arr)[j] > (*arr)[j+1] {
// 			// 交换
// 			temp = (*arr)[j]
// 			(*arr)[j] = (*arr)[j+1]
// 			(*arr)[j+1] = temp
// 		}
// 	}
// 	fmt.Println("第二次排序后arr=", (*arr))
// 	// 完成第三轮排序(外层1次)
// 	for j := 0; j < 2; j++ {
// 		if (*arr)[j] > (*arr)[j+1] {
// 			// 交换
// 			temp = (*arr)[j]
// 			(*arr)[j] = (*arr)[j+1]
// 			(*arr)[j+1] = temp
// 		}
// 	}
// 	fmt.Println("第三次排序后arr=", (*arr))

// 	// 完成第四轮排序(外层1次)
// 	for j := 0; j < 1; j++ {
// 		if (*arr)[j] > (*arr)[j+1] {
// 			// 交换
// 			temp = (*arr)[j]
// 			(*arr)[j] = (*arr)[j+1]
// 			(*arr)[j+1] = temp
// 		}
// 	}
// 	fmt.Println("第四次排序后arr=", (*arr))
// }
func BubblieSort(arr *[5]int) {
	fmt.Println("排序钱arr=", (*arr))
	temp := 0 // 临时变量（用户做交换）
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			// if (*arr)[j] > (*arr)[j+1] { // 从小到大排序
			if (*arr)[j] < (*arr)[j+1] { // 从大到小排序
				// 交换
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("第四次排序后arr=", (*arr))
}
func main() {
	// 第一数组
	arr := [5]int{24, 69, 80, 57, 13}
	// 将数组传递个一个函数完成排序
	BubblieSort(&arr)

	fmt.Println("main arr=", arr) // 有序 是有序的

}
