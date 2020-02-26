package process

import (
	"fmt"
	message "go_bas/go_code/chatroom/common/massage"
)

// 客户端腰围的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

//var CurUser model.CurUser //我们在用户登录成功后，完成对CurUser初始化
// 在客户端显示当前用户
func outputOnlineUser() {
	// 遍历一把 onlineUsers
	fmt.Println("当前在线用户列表")
	for id := range onlineUsers {
		// 如果不显示自己

		fmt.Println("用户id:\t", id)
	}
}

// 编写一个方法， 处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NOtifyUserStatusMes) {
	//user := &message.User{
	//	//	UserId:     notifyUserStatusMes.UserId,
	//	//	UserStatus: notifyUserStatusMes.Status,
	//	//}
	//	//onlineUsers[notifyUserStatusMes.UserId] = user
	// 优化后代码
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { // 原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
