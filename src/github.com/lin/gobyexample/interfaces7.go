package main

import "fmt"

type Animals interface {
	Speak() string
}

type Dog struct {

}

type Cat struct {

}

func (d *Dog) Speak() string {
	return "Wow"
}

func (c *Cat) Speak() string {
	return "Meow"
}

func main() {
	var d Dog
	var animals Animals = &d
	fmt.Println(animals.Speak())
}
