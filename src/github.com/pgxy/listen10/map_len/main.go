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
func main() {
	map_len()
}
