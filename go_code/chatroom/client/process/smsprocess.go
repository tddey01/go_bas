package process

import (
	"encoding/json"
	"fmt"
	"go_bas/go_code/chatroom/client/utils"
	message "go_bas/go_code/chatroom/common/massage"
)

type SmsProcess struct {
}

// 发送群聊消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	// 穿件一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	// 创建一个SmsMes 实例
	var smsMes message.SmsMes
	smsMes.Content = content // 内容
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	// 序列化 smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(smsMes) err=", err.Error())
		return
	}
	mes.Data = string(data)
	// 再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(smsMes) err=", err.Error())
		return
	}
	//  将mes序列化后的发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	//发送
	err = tf.WritedPkg(data)
	if err != nil {
		fmt.Println("SendGroupMes tf.WritedPkg(data) err=", err.Error())
		return
	}
	return
}
