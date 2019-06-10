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
	url := "http://10.2.86.249/sso/login.jsp?version=null&appid=app1000&token=app1000&success=http%3A%2F%2Fhome.hq.cmcc%2Fhome.php%2FHome%2FUser%2Flogin?tempid=app10005d3d24d2-497f-36eb-1426-b94b1a98d1fc-1560135826&apptempid=app10005d3d24d2-497f-36eb-1426-b94b1a98d1fc-1560135826"
	for i:=0;i<100;i++{
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






