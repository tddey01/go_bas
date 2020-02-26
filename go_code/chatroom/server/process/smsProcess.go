package process2

import (
	"fmt"
	"go_bas/go_code/chatroom/common/massage"
	"go_bas/go_code/chatroom/server/utils"
	"net"

	"encoding/json"
)

type SmsProcess struct {
	//..[暂时不需字段]
}

//写方法转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	//遍历服务器端的onlineUsers map[int]*UserProcess,
	//将消息转发取出.
	//取出mes的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}
func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn, //
	}
	err := tf.WritedPkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
	}
}
