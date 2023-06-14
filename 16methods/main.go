package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Golang")
	// no inhertiance in golang
	// no super or parent keyword in golang

	aditi := User{"Aditi", "adi@go.dev", true, 25}
	fmt.Println(aditi)
	fmt.Printf("Aditi details are: %+v\n", aditi)
	aditi.GetStatus()
	aditi.NewMail()
	fmt.Println(aditi)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active: ",u.Status)
}

func(u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is" ,u.Email)
}
