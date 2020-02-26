package process2

import (
	"encoding/json"
	"fmt"
	message "go_bas/go_code/chatroom/common/massage"
	"go_bas/go_code/chatroom/server/model"
	"go_bas/go_code/chatroom/server/utils"
	"net"
)

//
type UserProcess struct {
	// 字段
	Conn net.Conn
	// 增加一个字段 表示conn是哪个用户
	UserId int
}

//编写一个 函数serverProcessRegister函数  专门处理注册请求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	//1.先从mes 中取出 mes.Data ，并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库去完成注册.
	//1.使用model.MyUserDao 到redis去验证
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXTSTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXTSTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4. 将data 赋值给 resMes
	resMes.Data = string(data)

	//5. 对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//6. 发送data, 我们将其封装到writePkg函数
	//因为使用分层模式(mvc), 我们先创建一个Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritedPkg(data)
	return

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

	// 我们到数据库redis 数据库完成验证
	//1 使用model。MyUserDao 到redis 去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXSTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误...."
		}

	} else {
		loginResMes.Code = 200
		//这里，因为用户登录成功，我们就把该登录成功的用放入到userMgr中
		//将登录成功的用户的userId 赋给 this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//将当前在线用户的id 放入到loginResMes.UsersId
		//遍历 userMgr.onlineUsers
		for id := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登录成功")
	}

	//// 如果用户id = 100 密码=123456 认为是合法 否则不合法
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	// 合法
	//	loginResMes.Code = 200
	//
	//} else {
	//	// 不合法
	//	loginResMes.Code = 500 // 500 状态码  表示该用户不存在
	//	loginResMes.Error = "表示该用户不存在 请注册在使用..."
	//}
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
