// $Env:GOOS="linux";$Env:GOARCH="mipsle";go build .\main.go

package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	// "github.com/lin/my"
)

func Main2(){
	fmt.Printf("hello\n")
	// my.Test()
	// my.Myip()
	// fmt.Scanln()
	url := "http://www.baidu.com"
	for i:=0;i<10;i++{
		res, _ := http.Get(url)
		fmt.Printf("%d, %s\n",res.StatusCode,res.Proto)
		b, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		fmt.Printf("%s\n", b)
		fmt.Printf("No.%d finished\n", i)
	}
	// fmt.Scanln()
}



// type Integer int 

// func (a Integer) Less(b Integer) bool{
// 	return a<b
// }

// func (a *Integer) Add(b Integer) {
// 	*a+=b
// }

// func (a *Integer) Substract(b Integer) {
// 	*a -= b
// }

// var v8 func(a int) int


func main(){
	// var a Integer = 1
	// fmt.Println(a.Less(0))
	// a.Add(2)
	// fmt.Println(a)
	// a.Substract(4)
	// fmt.Println(a)
	Main2()
}






