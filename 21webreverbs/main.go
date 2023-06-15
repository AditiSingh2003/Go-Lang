package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)


func main() {
	fmt.Println("Welcome to Web Verb")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "https://www.google.com"

	response,err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)
	fmt.Println("Content length is:",response.ContentLength)

	var responseString strings.Builder
	body,_ := ioutil.ReadAll(response.Body)
	byteCount,_ := responseString.Write(body)
	fmt.Println("Response body:",responseString.String())
	fmt.Println("ByteCount is: ",byteCount)
	

}