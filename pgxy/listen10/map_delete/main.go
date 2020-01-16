package main

import "fmt"

func map_delete() {
	var a map[string]int
	fmt.Printf("a=%v\n", a)
	//a["st01"] = 1000
	if a == nil {
		a = make(map[string]int, 16)
		fmt.Printf("a=%v\n", a)
		a["stu01"] = 1000
		a["stu02"] = 2000
		a["stu03"] = 3000
		fmt.Printf("a=%#v\n", a)
		delete(a, "stu02") //单条删除
		fmt.Printf("a=%#v\n", a)

		for key, _ := range a {
			delete(a, key) // 全部删除
		}
		fmt.Printf("after delete a=%#v\n", a)
	}
	//map删除元素
}

func main() {
	map_delete()

}
