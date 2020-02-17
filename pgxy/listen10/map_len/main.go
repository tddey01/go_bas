package main

import "fmt"

func map_len() {
	a := map[string]int{
		"stud01": 1000,
		"stud02": 2000,
	}
	a["stud03"] = 3000
	fmt.Printf("length is  %d\n", len(a))
}

func map_import() {
	a := map[string]int{
		"stud01": 1000,
		"stud02": 2000,
	}
	a["stud03"] = 3000
	fmt.Printf("length is  %v\n", a)
	b := a
	b["stud04"] = 5000
	fmt.Printf("length is  %v\n", a)
}

func main() {
	// map_len() //map是引⽤用类型
	map_import()
}
