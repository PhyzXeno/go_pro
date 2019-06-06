/*
	如何初始化一个go项目：
	1.设置GOPATH，这个指的是包含src，bin，pkg的那个目录，用于作为Workspace。不能与GOROOT相混淆，GOROOT是GO可执行程序的目录。
	2.一般来说GOPATH都是名为“go”的文件夹，其中有src、bin、pkg
	3.go程序员一般把所有代码放在一起。所以src中可以设立这样的一个目录 github.com/lin/hello
	4.然后在hello目录中，便可创建hello.go
	package main
	import "fmt"
	func main(){
	fmt.Printf("hello world\n")
	}

	notes:
	1.package代表一个文件夹，这个文件夹中的所有文件开头都要写package name。一般name为这个文件夹的名字，对于可执行的package，里面的文件开头要写package main
	2.在github.com/lin/hello中执行 go install 可在bin中生成可执行文件。也可直接执行 go install github.com/lin/hello，生成可执行文件
	3.远程的package，可以通过 go get github.com/golang/example/hello 来获取，这样便会在src中建立相应文件夹，获取相应package，方便引用
	4.不同文件夹中函数的引用方式是
	在github.com/lin/hello/hello.go中使用 import "github.com/golang/example/stringutil" ，然后 stringutil.Reverse("hello") 便可使用
*/
package my


