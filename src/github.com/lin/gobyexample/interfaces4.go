package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {

}

func (d Dog) Speak() string {
	return "wowo!"
}

type Cat struct {

}

func (c *Cat) Speak() string  {
	return "Meow!"
}

type Lilla struct {

}

func (l *Lilla) Speak() string {
	return "?????"
}

func main()  {
	animals := []Animal{&Dog{}, new(Cat), &Lilla{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	//li := new(Lilla)
	//fmt.Println(li.Speak())
	l := Lilla{}
	fmt.Println(l.Speak())
}
