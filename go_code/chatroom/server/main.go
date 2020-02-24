package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	message "go_bas/go_code/chatroom/common/massage"
	"io"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	// 先定义一个切片
	buf := make([]byte, 8096)
	fmt.Println("等待读取客户端发送的数据....")
	//这个conn.Read 在conn 没有关闭的情况下， 才会阻塞
	// 如果客户端关闭了  conn 则就不会阻塞了
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read(buf[:4]) err= ",err )
		return
	}
	//fmt.Println("读到的buf长度 =",buf[:4])
	//根据读到的长度 [:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// 根据 pkgLen 读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("conn.Read body  error")
		fmt.Println("conn.Read(buf[:pkgLen]) err=", err)
		return
	}

	// 吧pkgLen 反序列化成 -> message.Message
	//技术就是一层窗户纸 &mes!!
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen],&mes) err=", err)

		return
	}

	return
}

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 这个需要及时延迟关闭 conn
	defer conn.Close()

	// 读客户端发送的信息
	for {
		// 这里我们将读取数据包，直接封装成一个函数 readpkg()， 返回Message，Err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端推出了 服务器端也正常退出")
				return
			} else {
				fmt.Println("readPkg(conn) err=", err)
				return
			}
		}

		fmt.Println("mes=", mes)
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

		}
		//一旦连接成功， 则启动一个携程和客户端保持通讯
		go process(conn)
	}
}
