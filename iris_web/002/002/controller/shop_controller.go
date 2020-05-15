package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/tddey01/go_bas/iris_web/002/002/service"
)

//商店功能模块控制结构体
type ShopController struct {
	//上下文对象
	Ctx     iris.Context
	Service service.ShopService
	Session *sessions.Session
}
