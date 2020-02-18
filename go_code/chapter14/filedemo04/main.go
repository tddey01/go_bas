package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个新文件，写入内容 5句 "hello, Gardon"
	//1 .打开文件 test.txt
	filePath := "./test.txt"
	//file, err := os.OpenFile(filePath, O_WRONLY|O_CREATE, 0666)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	// 要及时关闭句柄，防止内存泄漏
	defer file.Close()
	//准备写入5句 "hello Gardon"
	str := "hello,Gardon\r\n" // /n表示换行符 转义

	// 写入时使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为writer是自带缓存的，因此在调用WriteString方法时，其实内容先写入到缓存的,所以需要调用Flush这个方法，将缓存数据真正的写入到文件 数据落地磁盘,否则文件会丢失数据
	writer.Flush()
}
