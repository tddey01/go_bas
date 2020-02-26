package process

import (
	"encoding/json"
	"fmt"
	"go_bas/go_code/chatroom/client/utils"
	message "go_bas/go_code/chatroom/common/massage"
	"net"
	"os"
)

// 显示登录成功后的界面
func ShowMenu() {
	fmt.Println("-------恭喜---登录成功----------")
	fmt.Println("-------1 显示在线用户列表-------")
	fmt.Println("-------2 发送消息 --------------")
	fmt.Println("-------3 信息列表 --------------")
	fmt.Println("-------4 退出系统 --------------")
	fmt.Println("请选择（1-4）:")
	var key int
	var content string
	// 因为我们总会使用到SmsProcess实例， 因为我们将其定义在swtich外部
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息")
		fmt.Println("请输入你想对大家说什么?")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统 ")
		os.Exit(0) //
	default:
		fmt.Println("你的输入有误 请重新输入")
	}
}

// 和服务器端保持通讯
func ProcessserverMes(conn net.Conn) {
	// 创建一个trnasfer实例  不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
		}
		// 如果读取到消息 又是下一步处理逻辑
		switch mes.Type {
		case message.NOtifyUserStatusMesType: // 有人上线了
			// 1 、 取出 NotifyUserStatusMes
			var notifyUserStatusMes message.NOtifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)

			// 2 把这个用户信息 转态保存到客户端map中
			updateUserStatus(&notifyUserStatusMes)
		// 处理
		case message.SmsMesType: // 有人全发消息
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器端返回了未知消息类型")
		}
		//	如果读取到消息， 又是下一步处理逻辑
		fmt.Printf("mes=%v\n", mes)
	}
}
