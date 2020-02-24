package main

import (
	"fmt"
	"io"
	"net" // 做网络socket开发是，net包含有我们需要所有的方法核函数
)

func process(conn net.Conn) {
	// 这里我们循环的接收客户单发送的数据
	defer conn.Close() // 关闭conn

	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read()
		// 1 、 等待客户端通过conn发送信息给我
		// 2、 如果客户端没有Write[发送]， 那么携程就阻塞这里了
		// fmt.Println("服务器在等待客户端发送信息", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从conn 读取
		if err == io.EOF {
			fmt.Println("服务器的退出")
			return
		}
		// 3 、显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听了...")
	// 1 tcp 表示使用网络协议是tcp
	//  8888 表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888") // 这种方法支持ipv4  "127.0.0.1:8888"   支持ipv6 "0.0.0.0:8888"
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() // 延时关闭

	// 循环等待客户端来连接我
	for {
		// 等待客户端链接我
		fmt.Println("等待客户端链接我")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc conn=%v  客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		// 这里要准备起一个携程，为客户端服务
		go process(conn)
	}
	// fmt.Printf("listen suc=%v\n", listen)
}
