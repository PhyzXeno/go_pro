package main

// import "net"
import "fmt"
import "github.com/lin/mysocks"

func main(){
	// listenAddr := "127.0.1.1:8080"
	// structListenAddr, err := net.ResolveTCPAddr("tcp", listenAddr)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(structListenAddr.IP[0:])
	// fmt.Println(structListenAddr.Port)
	fmt.Println(lightsocks.RandPassword())
}