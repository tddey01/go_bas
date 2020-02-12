package main

<<<<<<< HEAD
import "fmt"

//type AInterface interface {
//	Test01()
//	Test02()
//}
//
//type  BInterface interface {
//	Test01()
//	Test02()
//}
//type CInterface interface {
//	AInterface
//	BInterface // 错误 重复的 函数 接口调用
//}

type Usb interface {
	Say()
}

=======
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
>>>>>>> d6eff17dd0316f0ba44c76563534eca5b4b3277b
type Stu struct {

}

<<<<<<< HEAD
func (this *Stu) Say(){
	fmt.Println("Say()")
}
func main(){

	var stu Stu = Stu{}
	// 错误！ 会报 Stu类型没有实现Usb接口 ,
	// 如果希望通过编译,  var u Usb = &stu
	var u Usb =  &stu
	u.Say()
	fmt.Println("here ", u)
}

=======
func (stu Stu) Test01(){

}
func (stu Stu) Test02(){}

func main(){

}
// 错误 重复的 函数 接口调用