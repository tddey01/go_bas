package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello,北京!"
	//1 .打开文件已经存在文件, d:/abc.txt
	filePath := "./test.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	defer file.Close()
	// 先读取原来的文件内容 显示终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //如果读取到文件的末尾
			break
		}
		//显示到终端
		fmt.Print(str)
	}

	str := "hello,北京!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
