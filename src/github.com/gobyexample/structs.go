package main

import "fmt"

type person struct {
	name string
	age int
}

func main()  {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{"Alice", 30})

	fmt.Println(person{"Fred", 0})

	fmt.Println(&person{"Ann", 40})

	s := person{"Sean", 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(s.age)

	var lin person
	lin = person{"LIN", 26}
	fmt.Println(lin)
}
