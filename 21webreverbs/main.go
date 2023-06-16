// check for the server it should be up and running and

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


func main() {
	fmt.Println("Welcome to Web Verb")
	fmt.Println("Performing Get Request")
	PerformGetRequest()
	fmt.Println("Performing Post Request")
	PerformPostJsonRequest()
	fmt.Println("Performing Post Form Request")
	PerformPostFormRequest()

}

func PerformGetRequest(){
	const myurl = " http://localhost:8000"

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

func PerformPostJsonRequest(){
	const myurl = " http://localhost:8000"
	// send to correct url
	// type of data tell
	// send data

	requestBody := strings.NewReader(`
	{
		"Name" : "Aditi Singh",
		"Age" : 21,
		"Gender ": 
		"Female"
	}`) //fake json data
		response, err := http.Post(myurl,"application/json",requestBody)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		content , err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		// fmt.Println("Response status:", response.Status)
		// fmt.Println("Content length is:",response.ContentLength)
		fmt.Println("Response body:",string(content))
		
}

func PerformPostFormRequest(){
	const myurl = " http://localhost:8000"

	// formdata

	data := url.Values{}
	data.Add("Name","Aditi Singh")
	data.Add("Age","21")
	data.Add("email","aditisingh@gmail.com")

	response, err := http.PostForm(myurl,data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content , err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response body:",string(content))
}

