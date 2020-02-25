package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	message "go_bas/go_code/chatroom/common/massage"
	"go_bas/go_code/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//暂时不需要字段...

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
	fmt.Printf("客户端 发送消息长度=%d 内容=%v\n", len(data), string(data))
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
