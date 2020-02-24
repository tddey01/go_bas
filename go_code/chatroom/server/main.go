package main

import (
	"fmt"
	"net"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 这个需要及时延迟关闭 conn
	defer conn.Close()
	buf := make([]byte, 8096)
	// 读客户端发送的信息
	for {
		// 先定义一个切片

		fmt.Println("等待读取客户端发送的数据....")
 		_,err := conn.Read(buf[:4])
 		if err != nil{
 			fmt.Println("conn.Read(buf[:4]) err=",err)
			return
		}
		fmt.Println("读到的buf长度 =",buf[:4])
	}

}

func main() {
	// 提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	// 一旦监听成功， 就等待客户连接服务器
	for {
		fmt.Println("等待客户端连接服务器....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err=", err)
			return
		}
		//一旦连接成功， 则启动一个携程和客户端保持通讯
		go process(conn)
	}
}
