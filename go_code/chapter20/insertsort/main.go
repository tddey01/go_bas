package main

import (
	"fmt"
)

func InsertSort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		// 完成第一次， 给第二个元素找到合适位置并插入
		insertVal := arr[i]
		insertIndex := i - 1 // 下标

		// 从大到小 排序
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入后%v\n", i, *arr)
	}

	/*
		// 完成第一次， 给第二个元素找到合适位置并插入
		insertVal := arr[1]
		insertIndex := 1 - 1 // 下标

		// 从大到小 排序
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != 1 {
			arr[insertIndex+1] = insertVal
		}
		fmt.Println("第一次插入后", *arr)

		// 完成第二次， 给第三个元素找到合适位置并插入
		insertVal = arr[2]
		insertIndex = 2 - 1 // 下标

		// 从大到小 排序
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != 2 {
			arr[insertIndex+1] = insertVal
		}
		fmt.Println("第二次插入后", *arr)

		// 完成第三次， 给第四个元素找到合适位置并插入
		insertVal = arr[3]
		insertIndex = 3 - 1 // 下标

		// 从大到小 排序
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != 3 {
			arr[insertIndex+1] = insertVal
		}
		fmt.Println("第三次插入后", *arr)

		// 完成第四次， 给第五个元素找到合适位置并插入
		insertVal = arr[4]
		insertIndex = 4 - 1 // 下标

		// 从大到小 排序
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != 4 {
			arr[insertIndex+1] = insertVal
		}
		fmt.Println("第四次插入后", *arr)
	*/
}

func main() {
	arr := [5]int{23, 0, 12, 56, 34}

	fmt.Println("原始 数组 =", arr)
	InsertSort(&arr)

	fmt.Println("main 函数=", arr)
	fmt.Println(arr)
}
