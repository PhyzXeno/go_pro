package main

import "fmt"

func main() {
	s := make([][]int, 5)
	for i:=0;i<5;i++ {
		innerLen := 5
		s[i] = make([]int, innerLen)
		for j:=0;j<5;j++{
			if i==j {
				s[i][j] = i
			}else {
				s[i][j] = 0
			}
		}
	}
	//fmt.Println("s:", s)
	//
	//for i:=0;i<5;i++{
	//	fmt.Println(s[i])
	//}
	//
	Len := 5
	s3 := make([][][]int, Len)
	for i:=0;i<Len;i++ {
		s3[i] = make([][]int, Len)
		for j:=0;j<Len;j++ {
			s3[i][j] = make([]int, Len)
			for k:=0;k<Len;k++ {
				if i==j && j==k {
					s3[i][j][k] = i
				}else {
					s3[i][j][k] = 0
				}
			}
		}
	}

	fmt.Println("s3:", s3)
	for i:=0;i<Len;i++ {
		fmt.Println(s3[i])
	}
	for i:=0;i<Len;i++ {
		for j:=0;j<Len;j++ {
			fmt.Println(s3[i][j])
		}
		fmt.Println()
	}
}
