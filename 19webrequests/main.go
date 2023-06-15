package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://chat.openai.com/"

func main(){
	fmt.Println("Welcome to web requests in Golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response) // this is a pointer to a response object
	// fmt.Println("Response is: ", response)
	defer response.Body.Close()
	databytes , err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("Content is: ", content)
}