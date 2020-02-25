package main

import (
	"fmt"
	message "go_bas/go_code/chatroom/common/massage"
	process2 "go_bas/go_code/chatroom/server/process"
	"go_bas/go_code/chatroom/server/utils"
	"io"
	"net"
)

// 先创建一个Processor 的结构体
type Processor struct {
	Conn net.Conn
}

//开始编写一个serverProcessMes 函数
// 功能：根据客户端发送消息种类不同， 决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	//
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录逻辑
		// 先创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
	//注册逻辑
	default:
		fmt.Println("消息类型不存在  无法处理.....")
	}
	return
}

func (this *Processor) process2() (err error) {
	// 读客户端发送的信息
	for {
		// 这里我们将读取数据包，直接封装成一个函数 readpkg()， 返回Message，Err
		// 创建一个Transfer 实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端推出了 服务器端也正常退出")
				return err
			} else {
				fmt.Println("readPkg(conn) err=", err)
				return err
			}
		}
		//fmt.Println("mes=", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("err=", err)
			return err
		}
	}
}
