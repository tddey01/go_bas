package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开一个存在的文件中，将原来的内容覆盖成新的内容10句 "你好，尚硅谷!"

	//创建一个新文件，写入内容 5句 "hello, Gardon"
	//1 .打开文件已经存在文件, d:/abc.txt
	filePath := "./test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	defer file.Close()

	str := "你好，尚硅谷!1\r\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}
