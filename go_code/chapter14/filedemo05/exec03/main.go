package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开一个存在的文件，在原来的内容追加内容 'ABC! ENGLISH!'
	//1 .打开文件已经存在文件, d:/abc.txt
	filePath := "./test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	defer file.Close()

	str := "ABC! ENGLISH!\r\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}
