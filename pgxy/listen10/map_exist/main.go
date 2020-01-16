package main

import (
	"fmt"
)

//如何判断map指定的key是否存在? value, ok := map[key]
func map_exist1() {
	var a map[string]int

	a = make(map[string]int, 16)
	a["stud01"] = 1000
	a["stud02"] = 2000
	a["stud03"] = 3000
	fmt.Printf("a =%#v\n", a)

	var result int
	var ok bool
	var  key string = "stud03"
	result,ok = a[key]
	if ok == false{
		fmt.Printf(" key %s is not  exist\n", key)
	}else{
		fmt.Printf("key %s is %d\n", key,result)
	}
	fmt.Printf("result:%d\n", result)
	//如何判断map指定的key是否存在? value, ok := map[key]
}

func main() {
	map_exist1()
}
