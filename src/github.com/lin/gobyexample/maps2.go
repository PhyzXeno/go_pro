package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["lin"] = 1
	m["can"] = 11

	value, prs := m["can"]
	fmt.Println(value, prs)

}
