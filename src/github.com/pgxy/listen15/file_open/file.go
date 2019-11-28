package main

import (
	"fmt"
	"os"
)

func main() {
	//  只读的方法打开
	file, err := os.Open("./file.go")
	if err != nil {
		fmt.Println("open file filed , err:", err)
		return
	}
	defer file.Close()
}
