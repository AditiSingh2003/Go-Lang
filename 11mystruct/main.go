package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Golang")
	// no inhertiance in golang
	// no super or parent keyword in golang

	aditi := User{"Aditi", "adi@go.dev", true, 25}
	fmt.Println(aditi)
	fmt.Printf("Aditi details are: %+v\n", aditi)
}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int
}
