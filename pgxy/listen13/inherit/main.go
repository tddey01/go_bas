package main

import "fmt"

type Animal struct {
	Name string
	Sex  string
}

func (a *Animal) Talk() {
	fmt.Printf("i` talk , i `m %s\n", a.Name)
}

type PurnAnimal struct {
}

func (p *PurnAnimal) Talk() {
	fmt.Println("buru dongwu  talk")
}

type Dog struct {
	Feet string
	// Animal
	*Animal
	*PurnAnimal
}

func (d *Dog) Eat() {
	fmt.Println("dog is eat")
}

//func (d *Dog) Talk() {
//	fmt.Println("dog is talking")
//}

func main() {
	var d *Dog = &Dog{
		Feet: "Four feet",
		Animal: &Animal{
			Name: "dog",
			Sex:  "xiong",
		},
	}
	d.Eat()
	//d.Talk()
	d.Animal.Talk() //多重继承与冲突解决
	d.PurnAnimal.Talk()
}

//匿匿名结构体与继承
