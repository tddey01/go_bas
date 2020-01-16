package main

import "fmt"

func modify(a map[string]int) {
	a["modify001"] = 10000
}

func map_copy() {
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
		b := a
		b["stud03"] = 4000
		fmt.Printf("after modify a:%v\n", a)
		modify(a)
		fmt.Printf("after modify a:%v\n", a)
	}
}

func main() {
	map_copy()

}
