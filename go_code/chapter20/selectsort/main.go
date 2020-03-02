package main

import "fmt"

// 编一个函数selectSort 完成排序
func SelectSort(arr *[6]int) {
	// 标准的访问方式
	//(*arr)[1] = 600
	//arr[1] = 600
	//	 1 先完成将第一个大值和arr[0] => 先易后难

	//	架设 arr[0] 最大值
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		// 2 遍历后面1---[len(arr) -1] 比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { // 找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//	交换
		if maxIndex != 0 {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}

		fmt.Printf("第%d次 %v\n", j+1, *arr)

	}

	/*	max := arr[0]
		maxIndex := 0
		// 2 遍历后面1---[len(arr) -1] 比较
		for i := 0 + 1; i < len(arr); i++ {
			if max < arr[i] { // 找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//	交换
		if maxIndex != 0 {
			arr[0], arr[maxIndex] = arr[maxIndex], arr[0]
		}

		fmt.Println("第1次循环", arr)
		max = arr[1]
		maxIndex = 1
		// 2 遍历后面 2---[len(arr) -1] 比较
		for i := 1; i < len(arr); i++ {
			if max < arr[i] { // 找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//	交换
		if maxIndex != 0 {
			arr[1], arr[maxIndex] = arr[maxIndex], arr[1]
		}

		fmt.Println("第2次循环", arr)

		max = arr[2]
		maxIndex = 2
		// 2 遍历后面 2---[len(arr) -1] 比较
		for i := 2 + 1; i < len(arr); i++ {
			if max < arr[i] { // 找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//	交换
		if maxIndex != 0 {
			arr[2], arr[maxIndex] = arr[maxIndex], arr[2]
		}

		fmt.Println("第3次循环", arr)

		max = arr[3]
		maxIndex = 3
		// 2 遍历后面 2---[len(arr) -1] 比较
		for i := 3 + 1; i < len(arr); i++ {
			if max < arr[i] { // 找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//	交换
		if maxIndex != 0 {
			arr[3], arr[maxIndex] = arr[maxIndex], arr[3]
		}

		fmt.Println("第4次循环", arr)
	*/
}

func main() {
	// 定义一个数组 从大到小
	arr := [6]int{10, 34, 19, 100, 80, 900}
	fmt.Println("原始=", arr)
	SelectSort(&arr)
	fmt.Println("main函数")
	fmt.Println(arr)
}
