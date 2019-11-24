package main

import "fmt"

type Animal struct {
	Name string
	Sex  string
}

func (a *Animal) Talk() {
	fmt.Printf("i` talk , i `m %s\n", a.Name)
}

type Dog struct {
	Feet string
	// Animal
	*Animal
}

func (d *Dog) Eat() {
	fmt.Println("dog is eat")
}

func (d *Dog) Talk() {
	fmt.Println("dog is talking")
}

func main() {
	var d *Dog = &Dog{
		Feet: "Four feet",
		Animal: &Animal{
			Name: "dog",
			Sex:  "xiong",
		},
	}
	d.Eat()
	d.Talk()
}

//匿匿名结构体与继承
