package main

import "fmt"

func main(){
	fmt.Println("Welcome to defer in Golang")
	defer println("deferred call")
	defer fmt.Println("deferred call2")
	defer fmt.Println("deferred call3")
	println("hello world")
	mydefer()
}

func mydefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
}

// Output:
// Welcome to defer in Golang
// hello world
// 4
// 3
// 2
// 1
// 0
// deferred call3
// deferred call2
// deferred call