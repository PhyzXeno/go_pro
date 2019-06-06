package main
import (
	"fmt"
	"github.com/lin/my"
	"net/http"
	"io/ioutil"
)

func main(){
	fmt.Printf("hello\n")
	my.Test()
	// my.Myip()
	// fmt.Scanln()
	url := "http://www.baidu.com"
	res, _ := http.Get(url)
	// fmt.Printf("%d, %s\n",res.StatusCode,res.Proto)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Printf("%s\n", body)
	fmt.Scanln()
}