package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	// fmt.Println("conn 成功=", conn)
	// 功能一， 客户端就可以发送单行数据，然后就推出
	reader := bufio.NewReader(os.Stdin) //os.stdin  代表标准输入【终端】
	// 从终端读取一行用户输入， 并准备发送给服务器
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}

		// 如果用户输入是 exit就是退出
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户单退出...")
			break
		}
		// 再将line 发送给 服务器
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
	}
	// fmt.Printf("客户端发送了 %d 字节的数据，并推出\n", n)
}
