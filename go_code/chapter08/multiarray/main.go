package main
import "fmt"


func main() {
	// 二维数组的演示
	/*
		0 0 0 0 0 0
		0 0 1 0 0 0
		0 2 0 3 0 0
		0 0 0 0 0 0
	*/
	// 定义 /声明  二维数组

	var arr [4][6]int
	// 赋初值
	arr[1][2] = 1
	arr[2][1] = 1
	arr[2][3] = 1
	
	//遍历二维数组 按照要求输出图形
	for i :=0;i<4;i++{
		// fmt.Println(arr[i])
		for  j :=0 ; j<6;j++{
			fmt.Print(arr[i][j], " ")
			// fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	// fmt.Println(arr)
}
