package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int{
	fmt.Println("processing pointer")
	return r.width * r.height
}

func (r rect) perim() int {
	fmt.Println("processing struct")
	return 2*(r.width + r.height)
}

func main()  {
	r := rect{10,5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

	//Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
}
