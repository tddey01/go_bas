package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_bas/go_code/chatroom/client/utils"
	message "go_bas/go_code/chatroom/common/massage"
	"net"
	"os"
)

type UserProcess struct {
	//暂时不需要字段...

}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	//  2 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3  创建一个Loginmes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//  4  将registerMes序列胡
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(LoginMes) err=", err)
		return
	}

	// 第五步  把data 赋给了mes.Data 字段
	mes.Data = string(data)
	// 6 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}

	// 创建一个Tranfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	// 发送给服务器端
	err = tf.WritedPkg(data)
	if err != nil {
		fmt.Println("注册发送信息出错了 err=", err)
	}

	mes, err = tf.ReadPkg() // mes 就是 RegisterMes
	if err != nil {
		fmt.Println("readPkg(conn) err", err)
		return
	}

	//将mes的Data 部分 反序列化成RegisterMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功 可以重新登录")
		// 这里退出
		os.Exit(0)
	} else {
		//} else if loginResMes.Code == 500 {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

// 写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	// 开始定协议...
	//fmt.Printf("userId=%v userPwd=%v\n", userId, userPwd)
	//	//return nil
	// 第一步  连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	//  2 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3  创建一个Loginmes 结构体
	var LoginMes message.LoginMes
	LoginMes.UserId = userId
	LoginMes.UserPwd = userPwd
	//  4  将loginMes序列胡
	data, err := json.Marshal(LoginMes)
	if err != nil {
		fmt.Println("json.Marshal(LoginMes) err=", err)
		return
	}
	// 第五步  把data 赋给了mes.Data 字段
	mes.Data = string(data)
	// 6 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}
	//7 到这个时候  ， 就是我们要发送的的消息
	//7.1 先把data的长度发送给服务器
	// 先获取到 data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// 先定义一个切片
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 现在发送这个长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}
	fmt.Printf("客户端 发送消息长度=%d 内容=%s\n", len(data), string(data))
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}
	// 休眠20
	//time.Sleep(time.Second * 10)
	//fmt.Println("休眠了10秒...")
	// 这里还要处理服务器返回消息
	// 创建一个Tranfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg() // mes 就是
	if err != nil {
		fmt.Println("readPkg(conn) err", err)
		return
	}
	//将mes的Data 部分 反序列化成LoginReMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//fmt.Println("登录成功")

		// 可以显示当前在线哦用户列表 遍历loginResMes。UsersId的切片
		fmt.Println("当前在线用户列表如下")
		for _, v := range loginResMes.UsersId {
			// 如果我们要求不显示自己在线，增加下面代码
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			// 完成 客户端的onlineUsers 初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user

		}
		fmt.Println("\n\n")

		// 这里我们还需要客户端启动一个携程
		// 该携程保持和服务器端的通信， 如果服务器有数据推送给客户端
		// 则接收并显示在客户端
		go ProcessserverMes(conn)

		//	 1、 显示我们登录成功的菜单[循环]
		for {
			ShowMenu()
		}
	} else {
		//} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
