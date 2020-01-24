package main

import "fmt"

func main() {
	// 使用常规的for循环切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	// slice := arr[1:4] // 20,30,40
	// slice := arr[0:len(arr)]
	slice := arr[:] // 简写 两头都是默认值
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}

	// 使用for-range  方式遍历
	for i, v := range slice {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	slice2 := slice[1:2] //  slice [ 20, 30, 40]    [30]
	slice2[0] = 100      // 因为arr , slice 和slice2 指向的数据空间是同一个，因此slice2[0]=100，其它的都变化

	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr=", arr)

	// 用append内置函数， 可以对切片进行动态追加
	var slice3 []int = []int{100, 200, 300}
	// 通过append直接给slice3追加具体的元素
	slice3 = append(slice3, 400, 500)
	// 通过append将切片slice3追加给slice3
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3=", slice3)

	// 切片可以copy 复制等操作
	var slice4 []int = []int{1, 2, 3, 4, 5, 6, 7}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4 =", slice4)
	fmt.Println("slice5 =", slice5)

}
