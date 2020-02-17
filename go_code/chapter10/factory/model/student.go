package model

// 定义一个结构体
type Student struct {
	Name  string
	Score float64
}

//因为 student结构体首字母小写， 因此是只能在model使用
//工厂模式

type student1 struct {
	Name  string
	score float64
}

//func (s *student1)
func NewStudent(n string, s float64) *student1 {
	return &student1{
		Name:  n,
		score: s,
	}
}

//如果score字段首字母小写，则，在其它包不可以直接方法，我们可以提供一个方法
func (s *student1) GetScore() float64 {
	return s.score // ok
}
