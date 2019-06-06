package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [][][]int  // this is a slice, a slice type has no specified length.
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)

	b := [...]string{"Penn", "Teller"}
	fmt.Println(b)
}
