package test

import (
	"reflect"
	"testing"
)

type user struct{
	Userld string
	Name string
}
func TestReflectStructPtr(t *testing.T){
	var {
		model *user
		sv  reflect.Type
		elem reflect.Value
	}
	st = reflect.TypeOf(model) // 获取类型 *user
	t.Log("reflect.TypeOf",st.Kind().String()) //ptr
	st = st.Elem() // st指向的类型
	t.Log("reflect.TypeOf.Elem",st.Kind().String()) // struct
	elem = reflect.New() // 返回一个Value类型值，改值持有一个指向类型为typ的新申请的零值的指针
	t.Log("reflect.New",elem.kind().String()) // ptr
	t.Log("reflec.New.Elem",elem.Elem().Kind().String()) //struct
	// model就是创建的user结构体变量（实例）
	model = elem.interface().(*user) // model 是*user他的指向的elem是一样的
	elem = elem.Elem() // 取得elem指向的值
	elem.FieldByName("Userld").SetString("12345678")
	elem.FieldByName("Name").SetString("nickname")
	t.Log("model model.Name",model,model.Name)
}