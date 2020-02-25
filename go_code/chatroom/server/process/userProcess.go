package process2

import (
	"encoding/json"
	"fmt"
	message "go_bas/go_code/chatroom/common/massage"
	"go_bas/go_code/chatroom/server/utils"
	"net"
)

//
type UserProcess struct {
	// 字段
	Conn net.Conn
}

// 编写一个函数serverProcessLogin函数， 专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	// 因为使用了 分层模式 (mvc) 我们先创建了一个Transfer实例， 然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritedPkg(data)
	return
}
