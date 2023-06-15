package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123"

func main(){
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println("myurl is: ", myurl)

	// Parse the URL

	result, err :=url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Result is of type: %T\n", result) //pointer
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	// Query string
	qparams := result.Query()
	fmt.Printf("Query params are of type: %T\n", qparams) //map
	fmt.Println("coursename is: ", qparams["coursename"])

	for key, val := range qparams {
		fmt.Println(key," Param is ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=aditi",
	}
	fmt.Println(partsOfUrl.RawPath)
	fmt.Println("Another URL is: ", partsOfUrl.String())

}