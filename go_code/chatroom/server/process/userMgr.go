package process2

import "fmt"

//因为UserMgr 实例 在服务器端且只有一个
// 因为在很多的地方， 都会使用 因此我们将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对userMgr初始化工作o

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对 onlineUsers 添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up // 添加
}

//删除
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前在线所有用户 查询
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// 根据id 返回一个对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	// 如何map取出一个值 带检查方式
	up, ok := this.onlineUsers[userId]
	if !ok { // 说明你要查找的这个用户 当前不在线
		err = fmt.Errorf("用户%d 不在线", userId)
		return
	}
	return
}
