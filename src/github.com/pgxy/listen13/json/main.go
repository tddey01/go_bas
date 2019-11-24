package main

import "fmt"

type Student struct {
	Id   string
	Name string
	Sex  string
}
type Class struct {
	Name    string
	Count   int
	Student []*Student
}

func main() {
	c := &Class{
		Name:  "101",
		Count: 0,
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%d", i),
			Sex:  "man",
			Id:   fmt.Sprintf("%d", i),
		}
		c.Student = append(c.Student, stu)
	}
}
