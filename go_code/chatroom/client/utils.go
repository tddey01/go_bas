package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	message "go_bas/go_code/chatroom/common/massage"
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
		fmt.Println("conn.Read(buf[:4]) err= ", err)
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

func writedPkg(conn net.Conn, data []byte) (err error) {
	// 先发一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(buf[:4]) err=", err)
		return
	}
	//发送data 数据本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	return
}
