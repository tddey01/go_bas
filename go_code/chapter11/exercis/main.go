package main

type AInterface interface {
	Test01()
	Test02()
}

type  BInterface interface {
	Test01()
	Test02()
}
type CInterface interface {
	AInterface
	BInterface
}
type Stu struct {

}

func (stu Stu) Test01(){

}
func (stu Stu) Test02(){}

func main(){

}
// 错误 重复的 函数 接口调用

