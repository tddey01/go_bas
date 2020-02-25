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
	if n != 4 || err != nil {
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

// 编写一个函数serverProcessLogin函数， 专门处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	// 核心代码
	//1 先从mes 中获取到mes.Data,并且直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginMes) err=", err)
		return
	}
	// 1 先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	// 2  在声明一个 LoginResMes 并完成赋值
	var loginResMes message.LoginResMes
	// 如果用户id = 100 密码=123456 认为是合法 否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		// 合法
		loginResMes.Code = 200

	} else {
		// 不合法
		loginResMes.Code = 500 // 500 状态码  表示该用户不存在
		loginResMes.Error = "表示该用户不存在 请注册在使用..."
	}
	// 3  将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) 序列化失败 err=", err)
	}
	// 4 将data 赋值给 resMes
	resMes.Data = string(data)
	// 5 对resMes 进行序列化 ， 准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err=", err)
	}
	// 6 发送data  我们将其封装到一个writePkg函数
	err = writedPkg(conn, data)
	return
}

//开始编写一个serverProcessMes 函数
// 功能：根据客户端发送消息种类不同， 决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	//
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录逻辑
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
	//注册逻辑
	default:
		fmt.Println("消息类型不存在  无法处理.....")
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
		//fmt.Println("mes=", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
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
