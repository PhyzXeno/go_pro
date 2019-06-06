package main

import (
	"fmt"
)

func sum(s []int, cha chan int) {
	total := 0
	for _,a := range s {
		total += a
	}
	cha <- total
}

func main() {
	nums := func () []int{
		s := make([]int, 100)
		for i:=0;i<100;i++ {
			s[i] = i
		}
		return s
	}()
	cha := make(chan int)
	go sum(nums[0:20], cha)
	go sum(nums[20:40], cha)
	go sum(nums[40:60], cha)
	go sum(nums[60:80], cha)
	go sum(nums[80:100], cha)
	t := make([]int, 5)
	total := 0
	for i:=0;i<5;i++ {
		t[i] = <-cha
		total += t[i]
	}
	fmt.Println(t, total)
}
