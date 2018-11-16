package main

import "fmt"
import "log"
import "net/http"
import "io/ioutil"
import "regexp"

func main() {
	url := "http://ip4.me"
	var str string
	var re = regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	str = string(body)
	fmt.Println("Current ip is:", re.FindString(str))
}
