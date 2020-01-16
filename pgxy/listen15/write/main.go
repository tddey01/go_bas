package main

import "os"

import "fmt"

import "bufio"

func main() {
	file, err := os.OpenFile("./test.dat", os.O_CREATE|os.O_WRONLY, 0666)
	// file, err := os.OpenFile("./test.dat", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed , err :", err)
		return
	}
	defer file.Close()

	// str := "hello world\n"
	// file.Write([]byte(str))
	// file.WriteString(str)
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello world\n")
		// 写入到缓存区域
	}
	writer.Flush()
}
