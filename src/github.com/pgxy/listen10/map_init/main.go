package main

import "fmt"

func main() {
	var a map[string]int = map[string]int{
		"stud01": 1000,
		"stud02": 2000,
		"stud03": 3000,
	}
	// 声明时进⾏行行初始化  
	fmt.Println(a)
	a["stud01"] = 8888
	a["stud04"] = 3333
	fmt.Println(a)


}
