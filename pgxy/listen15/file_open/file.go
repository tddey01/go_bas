package main

import (
	"fmt"
	"io"
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

	// var buf [128]byte
	// 	n, err := file.Read(buf[:])
	// 	fmt.Println("read length:", n)
	// 	fmt.Println("data:",string(buf[0:n]))
	var content []byte
	var buf [128]byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file:", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))

}
