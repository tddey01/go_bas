package model

import (
	message "go_bas/go_code/chatroom/common/massage"
	"net"
)

// 因为客户端， 我们很多地方就会使用到curlUser， 我们将其作为一个全局
type CurUser struct {
	Conn net.Conn
	message.User
}
