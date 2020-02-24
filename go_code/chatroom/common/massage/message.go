package message

// 确定消息类型 定义常量
const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"` // 消息的类型
	Data string `json:"data"` // 消息内容
}

// 先定义 两个消息 需要在增加
type LoginMes struct {
	UserId   int    `json:"userId`    // 用户id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

type LoginResMes struct {
	Code  int    `json:"code"`  // 返回状态吗    500 表示未注册用户  200 表示登录成功
	Error string `json:"error"` // 返回错误信息
}

type RegisterMes struct {
	//...
}
