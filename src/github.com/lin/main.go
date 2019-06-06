package main
import (
	"fmt"
	"net/http"
	// "io/ioutil"
	"github.com/lin/my"
)

func Main2(){
	fmt.Printf("hello\n")
	my.Test()
	// my.Myip()
	// fmt.Scanln()
	url := "http://www.baidu.com"
	res, _ := http.Get(url)
	// fmt.Printf("%d, %s\n",res.StatusCode,res.Proto)
	// b, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	// fmt.Printf("%s\n", body)
	// fmt.Scanln()
}



type Integer int 

func (a Integer) Less(b Integer) bool{
	return a<b
}

func (a *Integer) Add(b Integer) {
	*a+=b
}

func main(){
	var a Integer = 1
	fmt.Println(a.Less(0))
	a.Add(2)
	fmt.Println(a)
}


